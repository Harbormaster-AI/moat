import React, { Component } from 'react'
import LaboratoryOrderService from '../services/LaboratoryOrderService'

class ListLaboratoryOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                laboratoryOrders: []
        }
        this.addLaboratoryOrder = this.addLaboratoryOrder.bind(this);
        this.editLaboratoryOrder = this.editLaboratoryOrder.bind(this);
        this.deleteLaboratoryOrder = this.deleteLaboratoryOrder.bind(this);
    }

    deleteLaboratoryOrder(id){
        LaboratoryOrderService.deleteLaboratoryOrder(id).then( res => {
            this.setState({laboratoryOrders: this.state.laboratoryOrders.filter(laboratoryOrder => laboratoryOrder.laboratoryOrderId !== id)});
        });
    }
    viewLaboratoryOrder(id){
        this.props.history.push(`/view-laboratoryOrder/${id}`);
    }
    editLaboratoryOrder(id){
        this.props.history.push(`/add-laboratoryOrder/${id}`);
    }

    componentDidMount(){
        LaboratoryOrderService.getLaboratoryOrders().then((res) => {
            this.setState({ laboratoryOrders: res.data});
        });
    }

    addLaboratoryOrder(){
        this.props.history.push('/add-laboratoryOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">LaboratoryOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addLaboratoryOrder}> Add LaboratoryOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> TestCode </th>
                                    <th> FastingRequired </th>
                                    <th> SpecimenType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.laboratoryOrders.map(
                                        laboratoryOrder => 
                                        <tr key = {laboratoryOrder.laboratoryOrderId}>
                                             <td> { laboratoryOrder.testCode } </td>
                                             <td> { laboratoryOrder.fastingRequired } </td>
                                             <td> { laboratoryOrder.specimenType } </td>
                                             <td>
                                                 <button onClick={ () => this.editLaboratoryOrder(laboratoryOrder.laboratoryOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteLaboratoryOrder(laboratoryOrder.laboratoryOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewLaboratoryOrder(laboratoryOrder.laboratoryOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListLaboratoryOrderComponent
