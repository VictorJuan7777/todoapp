var btn = document.getElementById("login");
var datastore;
var subid;
btn.addEventListener("click",function(){
var username = document.getElementById("username").value;
var password = document.getElementById("password").value;
    var loginRequest = new XMLHttpRequest();
    loginRequest.open("POST","/login");
    loginRequest.setRequestHeader("Content-type","application/json");
    loginRequest.onload = function(){
        if (loginRequest.status == 200){
            var recData = JSON.parse(loginRequest.responseText);
            showActionForm("");
        } else {
            console.log(loginRequest.responseText)
        }
    };
    loginRequest.onerror = function(){
        console.log("Connection error");
    }
    loginRequest.send(
        JSON.stringify({
            "username" : username,
            "password" : password,
            })
    );
});

document.getElementById("logout").addEventListener("click",function(){
    var xhr = new XMLHttpRequest();
    xhr.open("POST","/logout");
    xhr.setRequestHeader("Content-type","application/json");
    xhr.onload = function(){
        
        if (xhr.status == 200){

            console.log("logout successful");
            document.getElementById("username").value="";
            document.getElementById("password").value="";
            document.getElementById("loginform").style.display="block";
            document.getElementById("logoutform").style.display="none";
            document.getElementById("json-content").innerHTML="";
        }
    };
    xhr.onerror = function(){
        console.log("Connection error");
    };
    xhr.send();
});

function checkCookie(){
    var xhr = new XMLHttpRequest();
    xhr.open("GET","/refresh");
    xhr.setRequestHeader("Content-type","application/json");
    xhr.onload = function(){
        
        if (xhr.status == 200){
            console.log("set cookie");
            showActionForm();
        }
    };
    xhr.onerror = function(){
        console.log("Connection error");
    };
    xhr.send();
};

function showActionForm(){
    document.getElementById("loginform").style.display="none";
    document.getElementById("logoutform").style.display="block";
    document.getElementById("showUser").innerHTML="Hello : "+getCookieValue("Username")+" !";
    var xhr = new XMLHttpRequest();
    xhr.open("GET","/action");
    xhr.setRequestHeader("Content-type","application/json");
    xhr.onload = function(){
        if (xhr.status == 200){
            document.getElementById("json-content").style.display="block";
            datastore = JSON.parse(xhr.responseText);
            showallaction("");
        }
    };
    xhr.onerror = function(){
        console.log("Connection error");
    };
    xhr.send();


};

function getCookieValue(cookieName) {
    let cookie = {};
    document.cookie.split(';').forEach(function(el) {
      let [key,value] = el.split('=');
      cookie[key.trim()] = value;
    })
    return cookie[cookieName];
  };

function reload(){
    document.getElementById("json-subaction").innerHTML = "";
  }

  function showallaction(actionid){
    var xhr = new XMLHttpRequest();
    xhr.open("GET","/action",true);
    xhr.setRequestHeader("Content-type","application/json");
    xhr.onload = function(){
        datastore = JSON.parse(xhr.responseText);
        var newContent=" ";
        newContent += "<div name='action'><input type='text' id='addaction' placeholder='Add Actions'/><input type='submit' id='sendaction' value=\"Send\" onclick='addaction()'/></div>";
        for (index in datastore){
            newContent +=  "<div name='action' onclick='loadsubaction(this.id)'id='"+ index +"'><div name='cbox'><input type='checkbox' id='cbox"+index+"' onclick='checkCOM("+index+")'";
            
            if (datastore[index].action.completed){
                newContent += " checked/>";
                newContent += "</div><div style='text-decoration:line-through;' id='actiontext'/>" + datastore[index].action.title +"</div></div>";
            }else{
                newContent += " />";
                newContent += "</div><div id='actiontext'/>" + datastore[index].action.title +"</div></div>";
            }
        }
        newContent += "<div onclick='reload()' id='white'></div>";
        document.getElementById("json-content").innerHTML = newContent;
        loadsubaction(actionid);
    };
    xhr.onerror = function(){
        console.log("Connection error");
    };
    xhr.send();
}

function loadsubaction(actionid){
    var newSubContent=" ";
    if (String(actionid) !=""){
        newSubContent +=  "<div name='action'><input type='text' id='showactiontext'  onkeydown='retitle("+ actionid +")' value='" + datastore[actionid].action.title  +"'/><input type='submit' name='deleteaction' id='"+ actionid+"' value='Delete' onclick='deleteaction(this.id)'/></div>";
        newSubContent += "<div name='action'><input type='text' id='addsubaction' placeholder='Add Actions'/><input type='submit' id='sendaction' value=\"Send\" onclick='addsubaction("+actionid+")'/></div>";
        if (String(actionid) !=""){
            for (index in datastore[actionid].subaction){
                newSubContent +=  "<div name='action' ><input type='checkbox'name='cbox' id='scbox"+datastore[actionid].subaction[index].id+"' onclick='checkSubCOM("+actionid+","+index+",this.id)'"
                if (datastore[actionid].subaction[index ].completed){
                    newSubContent +=  "checked/>";
                    newSubContent +=  "<input type='text' style='text-decoration:line-through;' onkeydown='resubtitle("+actionid +","+datastore[actionid].subaction[index].id+")' name='actiontext' id='s"+ datastore[actionid].subaction[index].id +"' value='"+ datastore[actionid].subaction[index].title +"'/><input type='submit' name='deletesub' value='X' onclick='deletesub("+actionid+","+datastore[actionid].subaction[index].id+")'/></div>";
                }else{
                    newSubContent +=  "/>";
                    newSubContent +=  "<input type='text' onkeydown='resubtitle("+actionid +","+datastore[actionid].subaction[index].id+")' name='actiontext' id='s"+ datastore[actionid].subaction[index].id +"' value='"+ datastore[actionid].subaction[index].title +"'/><input type='submit' name='deletesub' value='X' onclick='deletesub("+actionid+","+datastore[actionid].subaction[index].id+")'/></div>";
                }
            
            }
        }
    }
    document.getElementById("json-subaction").innerHTML = newSubContent;
}

function addaction(){
    var xhr = new XMLHttpRequest();
    xhr.open("POST","/action",true);
    xhr.setRequestHeader("Content-type","application/json");
    xhr.onload = function(){
        if (xhr.status == 200){
            console.log("Add action successful");
            document.getElementById("addaction").value="";
            
        }
        showActionForm();
    };
    xhr.onerror = function(){
        console.log("Connection error");
    };
    xhr.send(JSON.stringify({
        "username" : getCookieValue("Username"),
        "title" : document.getElementById("addaction").value,
        }));   
}

function addsubaction(id){
    var xhr = new XMLHttpRequest();
    xhr.open("POST","/subaction",true);
    xhr.setRequestHeader("Content-type","application/json");
    xhr.onload = function(){
        
        if (xhr.status == 200){
            console.log("Add subaction successful");
            document.getElementById("addsubaction").value="";
        }
        showallaction(id);
    };
    xhr.onerror = function(){
        console.log("Connection error");
    };
    xhr.send(JSON.stringify({
        "action_id" : String(datastore[id].action.id),
        "title" : document.getElementById("addsubaction").value,
        }));   
}

function deleteaction(id){
    var xhr = new XMLHttpRequest();
    xhr.open("DELETE","/action",true);
    xhr.setRequestHeader("Content-type","application/json");
    xhr.onload = function(){
        if (xhr.status == 200){
            console.log("Delete action successful");
            showallaction("");
        }
    };
    xhr.onerror = function(){
        console.log("Connection error");
    };
    xhr.send(JSON.stringify({
        "id" : String(datastore[id].action.id),
        }));
}
function retitle(id){
    if(window.event.keyCode === 13) {
        var xhr = new XMLHttpRequest();
        xhr.open("PUT","/action",true);
        xhr.setRequestHeader("Content-type","application/json");
        xhr.onload = function(){
            if (xhr.status == 200){
                showallaction(id);
            }
        };
        xhr.onerror = function(){
            console.log("Connection error");
        };
        xhr.send(JSON.stringify({
            "id" : String(datastore[id].action.id),
            "completed" : document.getElementById("cbox"+id).checked,
            "title" : document.getElementById("showactiontext").value,
            }));
    }
}

function resubtitle(actionID,subid){
    if(window.event.keyCode === 13) {
        var xhr = new XMLHttpRequest();
        xhr.open("PUT","/subaction",true);
        xhr.setRequestHeader("Content-type","application/json");
        xhr.onload = function(){
            if (xhr.status == 200){
                showallaction(String(actionID));
            }
        };
        xhr.onerror = function(){
            console.log("Connection error");
        };
        xhr.send(JSON.stringify({
            "id" : String(subid),
            "completed" : document.getElementById("scbox"+subid).checked,
            "title" : document.getElementById("s"+subid).value,
            }));
    }
}



function deletesub(actionID,subactionID){
    var xhr = new XMLHttpRequest();
    xhr.open("DELETE","/subaction",true);
    xhr.setRequestHeader("Content-type","application/json");
    xhr.onload = function(){
        if (xhr.status == 200){
            console.log("Delete action successful");
            showallaction(String(actionID));
        }
    };
    xhr.onerror = function(){
        console.log("Connection error");
    };
    xhr.send(JSON.stringify({
        "id" : String(subactionID),
        }));
}

function checkCOM(id){
    if (document.getElementById("cbox"+id).checked){
        var check = true
    }else{
        var check = false
    }
    var xhr = new XMLHttpRequest();
        xhr.open("PUT","/action",true);
        xhr.setRequestHeader("Content-type","application/json");
        xhr.onload = function(){
            if (xhr.status == 200){
                showallaction(String(id));
            }
        };
        xhr.onerror = function(){
            console.log("Connection error");
        };
        xhr.send(JSON.stringify({
            "id" : String(datastore[id].action.id),
            "completed" : check,
            "title" : String(datastore[id].action.title),
            }));
}

function checkSubCOM(actionID,subID,ID){
    if (document.getElementById(ID).checked){
        var check = true
    }else{
        var check = false
    }
    var xhr = new XMLHttpRequest();
        xhr.open("PUT","/subaction",true);
        xhr.setRequestHeader("Content-type","application/json");
        xhr.onload = function(){
            if (xhr.status == 200){
                showallaction(String(actionID))
            }
        };
        xhr.onerror = function(){
            console.log("Connection error");
        };
        xhr.send(JSON.stringify({
            "id" : String(datastore[actionID].subaction[subID].id),
            "completed" : check,
            "title" : String(datastore[actionID].subaction[subID].title),
            }));
}