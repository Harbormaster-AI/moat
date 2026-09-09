import React, { Component } from 'react'
import InspectionLotService from '../services/InspectionLotService'

class ListInspectionLotComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inspectionLots: []
        }
        this.addInspectionLot = this.addInspectionLot.bind(this);
        this.editInspectionLot = this.editInspectionLot.bind(this);
        this.deleteInspectionLot = this.deleteInspectionLot.bind(this);
    }

    deleteInspectionLot(id){
        InspectionLotService.deleteInspectionLot(id).then( res => {
            this.setState({inspectionLots: this.state.inspectionLots.filter(inspectionLot => inspectionLot.inspectionLotId !== id)});
        });
    }
    viewInspectionLot(id){
        this.props.history.push(`/view-inspectionLot/${id}`);
    }
    editInspectionLot(id){
        this.props.history.push(`/add-inspectionLot/${id}`);
    }

    componentDidMount(){
        InspectionLotService.getInspectionLots().then((res) => {
            this.setState({ inspectionLots: res.data});
        });
    }

    addInspectionLot(){
        this.props.history.push('/add-inspectionLot/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InspectionLot List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInspectionLot}> Add InspectionLot</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> LotNumber </th>
                                    <th> Quantity </th>
                                    <th> SampleSize </th>
                                    <th> CreatedOn </th>
                                    <th> InspectionType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.inspectionLots.map(
                                        inspectionLot => 
                                        <tr key = {inspectionLot.inspectionLotId}>
                                             <td> { inspectionLot.lotNumber } </td>
                                             <td> { inspectionLot.quantity } </td>
                                             <td> { inspectionLot.sampleSize } </td>
                                             <td> { inspectionLot.createdOn } </td>
                                             <td> { inspectionLot.inspectionType } </td>
                                             <td> { inspectionLot.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editInspectionLot(inspectionLot.inspectionLotId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInspectionLot(inspectionLot.inspectionLotId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInspectionLot(inspectionLot.inspectionLotId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInspectionLotComponent
