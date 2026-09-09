import React, { Component } from 'react'
import System_Service from '../services/System_Service'

class ListSystem_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
                system_s: []
        }
        this.addSystem_ = this.addSystem_.bind(this);
        this.editSystem_ = this.editSystem_.bind(this);
        this.deleteSystem_ = this.deleteSystem_.bind(this);
    }

    deleteSystem_(id){
        System_Service.deleteSystem_(id).then( res => {
            this.setState({system_s: this.state.system_s.filter(system_ => system_.system_Id !== id)});
        });
    }
    viewSystem_(id){
        this.props.history.push(`/view-system_/${id}`);
    }
    editSystem_(id){
        this.props.history.push(`/add-system_/${id}`);
    }

    componentDidMount(){
        System_Service.getSystem_s().then((res) => {
            this.setState({ system_s: res.data});
        });
    }

    addSystem_(){
        this.props.history.push('/add-system_/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">System_ List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSystem_}> Add System_</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> OwnerDepartment </th>
                                    <th> SystemType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.system_s.map(
                                        system_ => 
                                        <tr key = {system_.system_Id}>
                                             <td> { system_.name } </td>
                                             <td> { system_.ownerDepartment } </td>
                                             <td> { system_.systemType } </td>
                                             <td>
                                                 <button onClick={ () => this.editSystem_(system_.system_Id)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSystem_(system_.system_Id)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSystem_(system_.system_Id)} className="btn btn-outline-info btn-sm">View </button>
                                             </td>
                                        </tr>
                                    )
                                }
                            </tbody>
                        </table>

                 </div>

            </div>
        )
    }
}

export default ListSystem_Component
