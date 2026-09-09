import React, { Component } from 'react'
import ImagingOrderService from '../services/ImagingOrderService'

class ListImagingOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                imagingOrders: []
        }
        this.addImagingOrder = this.addImagingOrder.bind(this);
        this.editImagingOrder = this.editImagingOrder.bind(this);
        this.deleteImagingOrder = this.deleteImagingOrder.bind(this);
    }

    deleteImagingOrder(id){
        ImagingOrderService.deleteImagingOrder(id).then( res => {
            this.setState({imagingOrders: this.state.imagingOrders.filter(imagingOrder => imagingOrder.imagingOrderId !== id)});
        });
    }
    viewImagingOrder(id){
        this.props.history.push(`/view-imagingOrder/${id}`);
    }
    editImagingOrder(id){
        this.props.history.push(`/add-imagingOrder/${id}`);
    }

    componentDidMount(){
        ImagingOrderService.getImagingOrders().then((res) => {
            this.setState({ imagingOrders: res.data});
        });
    }

    addImagingOrder(){
        this.props.history.push('/add-imagingOrder/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ImagingOrder List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addImagingOrder}> Add ImagingOrder</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> BodySite </th>
                                    <th> Contrast </th>
                                    <th> Modality </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.imagingOrders.map(
                                        imagingOrder => 
                                        <tr key = {imagingOrder.imagingOrderId}>
                                             <td> { imagingOrder.bodySite } </td>
                                             <td> { imagingOrder.contrast } </td>
                                             <td> { imagingOrder.modality } </td>
                                             <td>
                                                 <button onClick={ () => this.editImagingOrder(imagingOrder.imagingOrderId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteImagingOrder(imagingOrder.imagingOrderId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewImagingOrder(imagingOrder.imagingOrderId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListImagingOrderComponent
