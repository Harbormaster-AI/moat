import React, { Component } from 'react'
import Component_Service from '../services/Component_Service'

class ListComponent_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
                component_s: []
        }
        this.addComponent_ = this.addComponent_.bind(this);
        this.editComponent_ = this.editComponent_.bind(this);
        this.deleteComponent_ = this.deleteComponent_.bind(this);
    }

    deleteComponent_(id){
        Component_Service.deleteComponent_(id).then( res => {
            this.setState({component_s: this.state.component_s.filter(component_ => component_.component_Id !== id)});
        });
    }
    viewComponent_(id){
        this.props.history.push(`/view-component_/${id}`);
    }
    editComponent_(id){
        this.props.history.push(`/add-component_/${id}`);
    }

    componentDidMount(){
        Component_Service.getComponent_s().then((res) => {
            this.setState({ component_s: res.data});
        });
    }

    addComponent_(){
        this.props.history.push('/add-component_/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Component_ List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addComponent_}> Add Component_</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> PartNumber </th>
                                    <th> Name </th>
                                    <th> ComponentCategory </th>
                                    <th> SerializationMethod </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.component_s.map(
                                        component_ => 
                                        <tr key = {component_.component_Id}>
                                             <td> { component_.partNumber } </td>
                                             <td> { component_.name } </td>
                                             <td> { component_.componentCategory } </td>
                                             <td> { component_.serializationMethod } </td>
                                             <td>
                                                 <button onClick={ () => this.editComponent_(component_.component_Id)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteComponent_(component_.component_Id)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewComponent_(component_.component_Id)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListComponent_Component
