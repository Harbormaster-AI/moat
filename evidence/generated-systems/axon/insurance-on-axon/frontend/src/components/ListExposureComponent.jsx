import React, { Component } from 'react'
import ExposureService from '../services/ExposureService'

class ListExposureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                exposures: []
        }
        this.addExposure = this.addExposure.bind(this);
        this.editExposure = this.editExposure.bind(this);
        this.deleteExposure = this.deleteExposure.bind(this);
    }

    deleteExposure(id){
        ExposureService.deleteExposure(id).then( res => {
            this.setState({exposures: this.state.exposures.filter(exposure => exposure.exposureId !== id)});
        });
    }
    viewExposure(id){
        this.props.history.push(`/view-exposure/${id}`);
    }
    editExposure(id){
        this.props.history.push(`/add-exposure/${id}`);
    }

    componentDidMount(){
        ExposureService.getExposures().then((res) => {
            this.setState({ exposures: res.data});
        });
    }

    addExposure(){
        this.props.history.push('/add-exposure/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">Exposure List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addExposure}> Add Exposure</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> ExposureType </th>
                                    <th> Status </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.exposures.map(
                                        exposure => 
                                        <tr key = {exposure.exposureId}>
                                             <td> { exposure.exposureType } </td>
                                             <td> { exposure.status } </td>
                                             <td>
                                                 <button onClick={ () => this.editExposure(exposure.exposureId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteExposure(exposure.exposureId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewExposure(exposure.exposureId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListExposureComponent
