import React, { Component } from 'react'
import RoleService from '../services/RoleService';

class CreateRoleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                responsibility: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeresponsibilityHandler = this.changeresponsibilityHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            RoleService.getRoleById(this.state.id).then( (res) =>{
                let role = res.data;
                this.setState({
                    name: role.name,
                    responsibility: role.responsibility
                });
            });
        }        
    }
    saveOrUpdateRole = (e) => {
        e.preventDefault();
        let role = {
                roleId: this.state.id,
                name: this.state.name,
                responsibility: this.state.responsibility
            };
        console.log('role => ' + JSON.stringify(role));

        // step 5
        if(this.state.id === '_add'){
            role.roleId=''
            RoleService.createRole(role).then(res =>{
                this.props.history.push('/roles');
            });
        }else{
            RoleService.updateRole(role).then( res => {
                this.props.history.push('/roles');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeresponsibilityHandler= (event) => {
        this.setState({responsibility: event.target.value});
    }

    cancel(){
        this.props.history.push('/roles');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Role</h3>
        }else{
            return <h3 className="text-center">Update Role</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> responsibility:&emsp; </label>
                                                <input placeholder="responsibility" name="responsibility" className="form-control" value={this.state.responsibility} onChange={this.changeresponsibilityHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateRole}>Save</button>
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

export default CreateRoleComponent
