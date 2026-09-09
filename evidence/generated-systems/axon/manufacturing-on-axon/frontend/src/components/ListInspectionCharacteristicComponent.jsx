import React, { Component } from 'react'
import InspectionCharacteristicService from '../services/InspectionCharacteristicService'

class ListInspectionCharacteristicComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                inspectionCharacteristics: []
        }
        this.addInspectionCharacteristic = this.addInspectionCharacteristic.bind(this);
        this.editInspectionCharacteristic = this.editInspectionCharacteristic.bind(this);
        this.deleteInspectionCharacteristic = this.deleteInspectionCharacteristic.bind(this);
    }

    deleteInspectionCharacteristic(id){
        InspectionCharacteristicService.deleteInspectionCharacteristic(id).then( res => {
            this.setState({inspectionCharacteristics: this.state.inspectionCharacteristics.filter(inspectionCharacteristic => inspectionCharacteristic.inspectionCharacteristicId !== id)});
        });
    }
    viewInspectionCharacteristic(id){
        this.props.history.push(`/view-inspectionCharacteristic/${id}`);
    }
    editInspectionCharacteristic(id){
        this.props.history.push(`/add-inspectionCharacteristic/${id}`);
    }

    componentDidMount(){
        InspectionCharacteristicService.getInspectionCharacteristics().then((res) => {
            this.setState({ inspectionCharacteristics: res.data});
        });
    }

    addInspectionCharacteristic(){
        this.props.history.push('/add-inspectionCharacteristic/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">InspectionCharacteristic List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addInspectionCharacteristic}> Add InspectionCharacteristic</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> CharacteristicCode </th>
                                    <th> Name </th>
                                    <th> LowerSpecLimit </th>
                                    <th> UpperSpecLimit </th>
                                    <th> Target </th>
                                    <th> MeasurementType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.inspectionCharacteristics.map(
                                        inspectionCharacteristic => 
                                        <tr key = {inspectionCharacteristic.inspectionCharacteristicId}>
                                             <td> { inspectionCharacteristic.characteristicCode } </td>
                                             <td> { inspectionCharacteristic.name } </td>
                                             <td> { inspectionCharacteristic.lowerSpecLimit } </td>
                                             <td> { inspectionCharacteristic.upperSpecLimit } </td>
                                             <td> { inspectionCharacteristic.target } </td>
                                             <td> { inspectionCharacteristic.measurementType } </td>
                                             <td>
                                                 <button onClick={ () => this.editInspectionCharacteristic(inspectionCharacteristic.inspectionCharacteristicId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteInspectionCharacteristic(inspectionCharacteristic.inspectionCharacteristicId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewInspectionCharacteristic(inspectionCharacteristic.inspectionCharacteristicId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListInspectionCharacteristicComponent
