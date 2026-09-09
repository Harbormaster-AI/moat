import React, { Component } from 'react'
import InsuredObjectService from '../services/InsuredObjectService'

class ListInsuredObjectComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                insuredObjects: []
        }
        this.addInsuredObject = this.addInsuredObject.bind(this);
        this.editInsuredObject = this.editInsuredObject.bind(this);
        this.deleteInsuredObject = this.deleteInsuredObject.bind(this);
    }

    deleteInsuredObject(id){
        InsuredObjectService.deleteInsuredObject(id).then( res => {
            this.setState({insuredObjects: this.state.insuredObjects.filter(insuredObject => insuredObject.insuredObjectId !== id)});
        });
    }
    viewInsuredObject(id){
        this.props.history.push(`/view-insuredObject/${id}`);
    }
    editInsuredObject(id){
        this.props.history.push(`/add-insuredObject/${id}`);
    }

    componentDidMount(){
        InsuredObjectService.getInsuredObjects().then((res) => {
            this.setState({ insuredObjects: res.data});
        });
    }

    addInsuredObject(){
        this.props.history.push('/add-insuredObject/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InsuredObject List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInsuredObject}> Add InsuredObject</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Description </th>
                                    <th> SerialOrId </th>
                                    <th> PrimaryAddress </th>
                                    <th> ObjectType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.insuredObjects.map(
                                        insuredObject => 
                                        <tr key = {insuredObject.insuredObjectId}>
                                             <td> { insuredObject.description } </td>
                                             <td> { insuredObject.serialOrId } </td>
                                             <td> { insuredObject.primaryAddress } </td>
                                             <td> { insuredObject.objectType } </td>
                                             <td>
                                                 <button onClick={ () => this.editInsuredObject(insuredObject.insuredObjectId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInsuredObject(insuredObject.insuredObjectId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInsuredObject(insuredObject.insuredObjectId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInsuredObjectComponent
