import React, { Component } from 'react'
import SerialNumberService from '../services/SerialNumberService'

class ListSerialNumberComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                serialNumbers: []
        }
        this.addSerialNumber = this.addSerialNumber.bind(this);
        this.editSerialNumber = this.editSerialNumber.bind(this);
        this.deleteSerialNumber = this.deleteSerialNumber.bind(this);
    }

    deleteSerialNumber(id){
        SerialNumberService.deleteSerialNumber(id).then( res => {
            this.setState({serialNumbers: this.state.serialNumbers.filter(serialNumber => serialNumber.serialNumberId !== id)});
        });
    }
    viewSerialNumber(id){
        this.props.history.push(`/view-serialNumber/${id}`);
    }
    editSerialNumber(id){
        this.props.history.push(`/add-serialNumber/${id}`);
    }

    componentDidMount(){
        SerialNumberService.getSerialNumbers().then((res) => {
            this.setState({ serialNumbers: res.data});
        });
    }

    addSerialNumber(){
        this.props.history.push('/add-serialNumber/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SerialNumber List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSerialNumber}> Add SerialNumber</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Serial </th>
                                    <th> ActivationDate </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.serialNumbers.map(
                                        serialNumber => 
                                        <tr key = {serialNumber.serialNumberId}>
                                             <td> { serialNumber.serial } </td>
                                             <td> { serialNumber.activationDate } </td>
                                             <td> { serialNumber.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editSerialNumber(serialNumber.serialNumberId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSerialNumber(serialNumber.serialNumberId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSerialNumber(serialNumber.serialNumberId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSerialNumberComponent
