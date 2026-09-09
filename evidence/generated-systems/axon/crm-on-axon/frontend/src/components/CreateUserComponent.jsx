import React, { Component } from 'react'
import UserService from '../services/UserService';

class CreateUserComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                username: '',
                fullName: '',
                email: '',
                locale: '',
                role: '',
                status: ''
        }
        this.changeusernameHandler = this.changeusernameHandler.bind(this);
        this.changefullNameHandler = this.changefullNameHandler.bind(this);
        this.changeemailHandler = this.changeemailHandler.bind(this);
        this.changelocaleHandler = this.changelocaleHandler.bind(this);
        this.changeRoleHandler = this.changeRoleHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            UserService.getUserById(this.state.id).then( (res) =>{
                let user = res.data;
                this.setState({
                    username: user.username,
                    fullName: user.fullName,
                    email: user.email,
                    locale: user.locale,
                    role: user.role,
                    status: user.status
                });
            });
        }        
    }
    saveOrUpdateUser = (e) => {
        e.preventDefault();
        let user = {
                userId: this.state.id,
                username: this.state.username,
                fullName: this.state.fullName,
                email: this.state.email,
                locale: this.state.locale,
                role: this.state.role,
                status: this.state.status
            };
        console.log('user => ' + JSON.stringify(user));

        // step 5
        if(this.state.id === '_add'){
            user.userId=''
            UserService.createUser(user).then(res =>{
                this.props.history.push('/users');
            });
        }else{
            UserService.updateUser(user).then( res => {
                this.props.history.push('/users');
            });
        }
    }
    
    changeusernameHandler= (event) => {
        this.setState({username: event.target.value});
    }
    changefullNameHandler= (event) => {
        this.setState({fullName: event.target.value});
    }
    changeemailHandler= (event) => {
        this.setState({email: event.target.value});
    }
    changelocaleHandler= (event) => {
        this.setState({locale: event.target.value});
    }
    changeRoleHandler= (event) => {
        this.setState({role: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/users');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add User</h3>
        }else{
            return <h3 className="text-center">Update User</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> username:&emsp; </label>
                                                <input placeholder="username" name="username" className="form-control" value={this.state.username} onChange={this.changeusernameHandler}/>

                                            <label> fullName:&emsp; </label>
                                                <input placeholder="fullName" name="fullName" className="form-control" value={this.state.fullName} onChange={this.changefullNameHandler}/>

                                            <label> email:&emsp; </label>
                                                <input placeholder="email" name="email" className="form-control" value={this.state.email} onChange={this.changeemailHandler}/>

                                            <label> locale:&emsp; </label>
                                                <input placeholder="locale" name="locale" className="form-control" value={this.state.locale} onChange={this.changelocaleHandler}/>

                                            <label> Role:&emsp; </label>
                                                <select value={this.state.role} onChange={this.changeRoleHandler}>
                      <option name="Role" className="form-control" >
                          SalesRep
                      </option>
                      <option name="Role" className="form-control" >
                          SalesManager
                      </option>
                      <option name="Role" className="form-control" >
                          ServiceAgent
                      </option>
                      <option name="Role" className="form-control" >
                          MarketingSpecialist
                      </option>
                      <option name="Role" className="form-control" >
                          Administrator
                      </option>
                      <option name="Role" className="form-control" >
                          Executive
                      </option>
                    </select>

                                            <label> Status:&emsp; </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Active
                      </option>
                      <option name="Status" className="form-control" >
                          Inactive
                      </option>
                      <option name="Status" className="form-control" >
                          Locked
                      </option>
                      <option name="Status" className="form-control" >
                          PendingInvite
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateUser}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>
                   </div>
            </div>
        )
    }
}

export default CreateUserComponent
