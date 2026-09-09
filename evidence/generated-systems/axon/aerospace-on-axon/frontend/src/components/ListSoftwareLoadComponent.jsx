import React, { Component } from 'react'
import SoftwareLoadService from '../services/SoftwareLoadService'

class ListSoftwareLoadComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                softwareLoads: []
        }
        this.addSoftwareLoad = this.addSoftwareLoad.bind(this);
        this.editSoftwareLoad = this.editSoftwareLoad.bind(this);
        this.deleteSoftwareLoad = this.deleteSoftwareLoad.bind(this);
    }

    deleteSoftwareLoad(id){
        SoftwareLoadService.deleteSoftwareLoad(id).then( res => {
            this.setState({softwareLoads: this.state.softwareLoads.filter(softwareLoad => softwareLoad.softwareLoadId !== id)});
        });
    }
    viewSoftwareLoad(id){
        this.props.history.push(`/view-softwareLoad/${id}`);
    }
    editSoftwareLoad(id){
        this.props.history.push(`/add-softwareLoad/${id}`);
    }

    componentDidMount(){
        SoftwareLoadService.getSoftwareLoads().then((res) => {
            this.setState({ softwareLoads: res.data});
        });
    }

    addSoftwareLoad(){
        this.props.history.push('/add-softwareLoad/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">SoftwareLoad List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addSoftwareLoad}> Add SoftwareLoad</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Version </th>
                                    <th> LoadType </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.softwareLoads.map(
                                        softwareLoad => 
                                        <tr key = {softwareLoad.softwareLoadId}>
                                             <td> { softwareLoad.version } </td>
                                             <td> { softwareLoad.loadType } </td>
                                             <td>
                                                 <button onClick={ () => this.editSoftwareLoad(softwareLoad.softwareLoadId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteSoftwareLoad(softwareLoad.softwareLoadId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewSoftwareLoad(softwareLoad.softwareLoadId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListSoftwareLoadComponent
