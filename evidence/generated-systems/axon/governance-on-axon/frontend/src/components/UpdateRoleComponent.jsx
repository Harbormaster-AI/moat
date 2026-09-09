import React, { Component } from 'react'
import RoleService from '../services/RoleService';

class UpdateRoleComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                responsibility: ''
        }
        this.updateRole = this.updateRole.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeresponsibilityHandler = this.changeresponsibilityHandler.bind(this);
    }

    componentDidMount(){
        RoleService.getRoleById(this.state.id).then( (res) =>{
            let role = res.data;
            this.setState({
                name: role.name,
                responsibility: role.responsibility
            });
        });
    }

    updateRole = (e) => {
        e.preventDefault();
        let role = {
            roleId: this.state.id,
            name: this.state.name,
            responsibility: this.state.responsibility
        };
        console.log('role => ' + JSON.stringify(role));
        console.log('id => ' + JSON.stringify(this.state.id));
        RoleService.updateRole(role).then( res => {
            this.props.history.push('/roles');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Role</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> responsibility: </label>
                                                <input placeholder="responsibility" name="responsibility" className="form-control" value={this.state.responsibility} onChange={this.changeresponsibilityHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateRole}>Save</button>
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

export default UpdateRoleComponent
