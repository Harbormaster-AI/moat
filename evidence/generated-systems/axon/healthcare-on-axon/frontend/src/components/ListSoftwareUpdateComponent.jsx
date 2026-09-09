import React, { Component } from 'react'
import SoftwareUpdateService from '../services/SoftwareUpdateService'

class ListSoftwareUpdateComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                softwareUpdates: []
        }
        this.addSoftwareUpdate = this.addSoftwareUpdate.bind(this);
        this.editSoftwareUpdate = this.editSoftwareUpdate.bind(this);
        this.deleteSoftwareUpdate = this.deleteSoftwareUpdate.bind(this);
    }

    deleteSoftwareUpdate(id){
        SoftwareUpdateService.deleteSoftwareUpdate(id).then( res => {
            this.setState({softwareUpdates: this.state.softwareUpdates.filter(softwareUpdate => softwareUpdate.softwareUpdateId !== id)});
        });
    }
    viewSoftwareUpdate(id){
        this.props.history.push(`/view-softwareUpdate/${id}`);
    }
    editSoftwareUpdate(id){
        this.props.history.push(`/add-softwareUpdate/${id}`);
    }

    componentDidMount(){
        SoftwareUpdateService.getSoftwareUpdates().then((res) => {
            this.setState({ softwareUpdates: res.data});
        });
    }

    addSoftwareUpdate(){
        this.props.history.push('/add-softwareUpdate/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SoftwareUpdate List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSoftwareUpdate}> Add SoftwareUpdate</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Version </th>
                                    <th> AppliedDate </th>
                                    <th> UpdateType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.softwareUpdates.map(
                                        softwareUpdate => 
                                        <tr key = {softwareUpdate.softwareUpdateId}>
                                             <td> { softwareUpdate.version } </td>
                                             <td> { softwareUpdate.appliedDate } </td>
                                             <td> { softwareUpdate.updateType } </td>
                                             <td>
                                                 <button onClick={ () => this.editSoftwareUpdate(softwareUpdate.softwareUpdateId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSoftwareUpdate(softwareUpdate.softwareUpdateId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSoftwareUpdate(softwareUpdate.softwareUpdateId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSoftwareUpdateComponent
