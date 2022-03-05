<template>
  <div id="app">

    <!-- Rahul -->
    <h1>My User API</h1>

    <button v-show = "!isShow" @click="isShow = !isShow, getUsers()"> Show Users </button>

    <div v-show="isShow"> 

        <div class = "userTable">

            <div class = "tableId"> <b>         User Id     </b> </div>
            <div class = "tableFirstName"> <b>  First Name  </b> </div>
            <div class = "tableMiddleName"> <b> Middle Name </b> </div>
            <div class = "tableLastName"> <b>   Last Name   </b> </div>
            <div class = "tableActions"> <b>    Actions     </b> </div>

        </div>

        <div v-for="i in newVar" :key = i.u_id >

            <div class = "userTable">

                <div class="tableId">           {{i.u_id}}          </div>
                <div class="tableFirstName">    {{i.first_name}}    </div>
                <div class="tableMiddleName">   {{i.middle_name}}   </div>
                <div class="tableLastName">     {{i.last_name}}     </div>

                <div class="tableActions">

                    <button @click="getUserAddresses(i), hideElements(), j=i"> Show Address </button>
                            &emsp;
                    <button @click="setEditValues(i), showEditForm(), hideElements('edit')"> Edit </button>
                            &emsp;
                    <button @click="deleteUser(i), hideElements()"> Delete User </button>
                
                </div>

            </div>

            <div v-if="switchForAddressDisplay & i===j" class = "addressTable" >

                <div v-for="j in newVarAddresses" :key = j.a_id class = "addressSpecific" >
                    
                    <br>Street:     {{j.street}}
                    <br>Area:       {{j.area}} 
                    <br>Pincode:    {{j.pincode}} 
                    <br>City:       {{j.city}}
                    <!-- &emsp; -->
                
                </div>

            </div>

        </div>

        <div class = "formStyle" v-show="switchForNewForm">

            <form>
                Enter Details: <br>

                <input v-model="userId" type="text" id="id" name="id" placeholder='User Id'><br>
                <input v-model="firstName" type="text" id="fname" name="fname" placeholder='First Name'><br>
                <input v-model="middleName" type="text" id="mname" name="mname" placeholder='Middle Name'><br>
                <input v-model="lastName" type="text" id="lname" name="lname" placeholder='Last name'><br>

                <button @click="newUser(), preventEvent($event)"> Create </button>
                &emsp;
                <button @click="hideElements(), preventEvent($event)"> Cancel </button>

            </form>

        </div>
          
        <div class = "formStyle" v-show="switchForEditForm"> 

            <form> 
                Edit details: <br> 

                <input v-model="firstName" type="text" id="fname" name="fname"><br>
                <input v-model="middleName" type="text" id="mname" name="mname"><br>
                <input v-model="lastName" type="text" id="lname" name="lname"><br>

                <button @click="preventEvent($event), editUserId()"> Update </button>
                        &emsp;
                <button @click="preventEvent($event), hideElements()"> Cancel </button>

            </form>

        </div>

        <div>
            <button @click="showNewForm(), hideElements('newForm')" v-show="!switchForNewForm"> New User </button>
        </div>

    </div>

    <!-- THIS IS DEDICATED DELETE MESSAGE SHOWING DIV -->
    <div v-show="switchForAPIResponse">   <p> {{reqResponse}} </p>    </div>

  </div>
</template>
// 🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹

<script>
import axios from 'axios';

export default {
    name: "App",

    data() {
        return {
        j:[],
        newVar: [],
        specificUser: [],
        newVarAddresses: [],
        switchForAPIResponse: false,
        isShow:false,
        deleteSwitch:false,
        switchForNewForm:false,
        switchForEditForm:false,
        switchForAddressDisplay:false,
        userId: "",
        lastName: "",
        firstName: "",
        middleName: "",
        reqResponse: "",
        userIdForGet: "",
        userIdForEdit: "",
        newVarFirstName: "",
        increment: Number,
        };
    },

    mounted() {
        console.log();
    },
// console.log();
    methods: {

        preventEvent(event) {
            event.preventDefault();
        },

        showNewForm() {
            this.userId = "";
            this.firstName = "";
            this.middleName = "";
            this.lastName = "";
            this.switchForNewForm = true;
        },

        showEditForm() {
            this.switchForEditForm = true;
        },

        showAddForm() {
        this.switchForAddressDisplay = true;
        },

        hideElements(keep) {
            // console.log(keep);
            this.switchForNewForm = false;
            this.switchForEditForm = false;
            this.switchForAddressDisplay = false;
            if (keep === "edit") {
                this.switchForEditForm = true;
                // switchForEditForm
            } else if (keep === "newForm") {
                this.switchForNewForm = true;
            }
        },

        setShow () {
            this.switchForAPIResponse = true;
            setTimeout(() => {
                this.switchForAPIResponse = false;
                this.reqResponse = '';
            }, 1500);
        },

        async getUsers() {
            try {
                const headers = { "Content-Type": "application/json" };
                const response = await axios.get("http://localhost:5435/v1/users/", headers);
                this.newVar = response.data.Data
            } catch(err) {
                console.log(err);
            } finally {
                // anything I want to execute regardless of try/catch condition
            }
        },

        async getUserAddresses (i) {
            this.newVarAddresses = [];
            try {
                const headers = { "Content-Type": "application/json" };
                let url = "http://localhost:5435/v1/users/" + i.u_id + "/addresses/"
                const response = await axios.get(url, headers); 
                this.newVarAddresses = response.data.Data
            } catch (err) {
                console.log(err);
            } finally {
                if (this.newVarAddresses.length === 0 ) {
                    this.reqResponse = "No Addresses";
                } else {
                    this.showAddForm();
                }
                this.setShow();
            }
        }, 

        setEditValues(i) {
            this.userIdForEdit = i.u_id;
            this.firstName = i.first_name;
            this.middleName = i.middle_name;
            this.lastName = i.last_name;
        },

        async editUserId () {
            if ( this.validateNames ( this.firstName, this.middleName, this.lastName )) {
                try {
                    const headers = { "Content-Type": "application/json" };
                    let url = 'http://localhost:5435/v1/users/' + this.userIdForEdit + '/';
                    let jsonRes = {
                        "first_name": this.firstName, 
                        "middle_name": this.middleName, 
                        "last_name": this.lastName,
                    }
                    const response = await axios.put(url, jsonRes, headers);                     
                    this.reqResponse = response.data.Message;
                } catch (err) {
                    console.log(err);
                } finally {
                    this.hideElements();
                }
                this.getUsers();
            }
            this.setShow();
        }, 

// newUser is now fully functional, except - the duplicate entry error is not yet visibile due to some issues
// so that, is to be solved...
        async newUser () {
            if ( this.validateNames ( this.firstName, this.middleName, this.lastName )) {
                try {
                    const headers = { "Content-Type": "application/json" };
                    let url
                    if (this.userId === '' ) {
                        url = 'http://localhost:5435/v1/users/'
                    } else {
                        url = 'http://localhost:5435/v1/users/' + this.userId + '/'
                    }

                    let jsonRes = {
                        "first_name": this.firstName,
                        "middle_name": this.middleName,
                        "last_name": this.lastName,
                    }
                    const response = await axios.post(url, jsonRes, headers);
                    this.reqResponse = response.data.Message;
                } catch (err) {
                    console.log(err);
                } finally {
                    // this.hideElements('newForm');
                    this.userId = "";
                    this.firstName = "";
                    this.middleName = "";
                    this.lastName = "";
                    this.hideElements();
                }
                this.getUsers();
            }
            this.setShow();
        },

        async deleteUser(i) {
            try {
                console.log(i.u_id);
                const headers = { "Content-Type": "application/json" };
                let url = 'http://localhost:5435/v1/users/' + i.u_id + '/'
                const response = await axios.delete(url, headers);
                this.reqResponse = response.data.Message;
            } catch (err) {
                console.log(err);                
            } finally {
                this.hideElements();
                this.setShow();
                this.getUsers();
            }
        },

        validateNames(first, middle, last) {

            let x = /[^A-Za-z]/
            if ( x.test(first) || x.test(middle) || x.test(last)){
                this.reqResponse = "Names cannot contain Symbols/Numbers";
                return false;

            } else if ( first == "" || middle == "" || last == ""){
                this.reqResponse = "Names cannot be empty.";
                return false;

            } else {
                return true;
            }

        },

    },
};

</script>

// 🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹🔹
<style>

#app {
  font-family: "Courier New", Courier, monospace;
  text-align: center;
  color: #2C3E50;
  margin-right: 360px;
  margin-left: 360px;
}

.addressTable {
  display: flex;
}

.addressSpecific {
  margin-bottom: 20px;
  border: 1px dashed #595959;
  width: fit-content;
  padding: 5px;
  text-align: left;
  display: flex;
  width: fit-content;
  margin-right: 20px;
}

.userTable {
  display: flex;
  margin-bottom: 20px;
  border: 1px dashed #595959;
  padding: 8px;
  align-items: center;
}

.tableId {
  display: flex;
  width: 10%;
  margin-left: 10px;
}

.tableFirstName {
  display: flex;
  width: 15%;
}

.tableMiddleName {
  display: flex;
  width: 15%;
}

.tableLastName {
  display: flex;
  width: 30%;
}

.tableActions {
  display: flex;
  width: 30%;
}

.formStyle {
  display: flex;
  width: 15%;
  font-size: medium;
  margin-bottom: 20px;
  border: 1px dashed #595959;
  padding: 10px;
  line-height: 25px;
}

button{
  font-family: "Courier New", Courier, monospace;
}

input{
    font-family: "Courier New", Courier, monospace;
}

::placeholder{
  font-family: "Courier New", Courier, monospace;
}

</style>
