import React, { Component } from 'react'
import ImagingCenterService from '../services/ImagingCenterService'

class ListImagingCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
                imagingCenters: []
        }
        this.addImagingCenter = this.addImagingCenter.bind(this);
        this.editImagingCenter = this.editImagingCenter.bind(this);
        this.deleteImagingCenter = this.deleteImagingCenter.bind(this);
    }

    deleteImagingCenter(id){
        ImagingCenterService.deleteImagingCenter(id).then( res => {
            this.setState({imagingCenters: this.state.imagingCenters.filter(imagingCenter => imagingCenter.imagingCenterId !== id)});
        });
    }
    viewImagingCenter(id){
        this.props.history.push(`/view-imagingCenter/${id}`);
    }
    editImagingCenter(id){
        this.props.history.push(`/add-imagingCenter/${id}`);
    }

    componentDidMount(){
        ImagingCenterService.getImagingCenters().then((res) => {
            this.setState({ imagingCenters: res.data});
        });
    }

    addImagingCenter(){
        this.props.history.push('/add-imagingCenter/_add');
    }

    render() {
        return (
            <div>
                 <h2 className="text-center">ImagingCenter List</h2>
                 <div className = "row">
                    <button className="btn btn-primary btn-sm" onClick={this.addImagingCenter}> Add ImagingCenter</button>
                 </div>
                 <br></br>
                 <div className = "row">
                        <table className = "table table-striped table-bordered">

                            <thead>
                                <tr>
                                    <th> Name </th>
                                    <th> Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    this.state.imagingCenters.map(
                                        imagingCenter => 
                                        <tr key = {imagingCenter.imagingCenterId}>
                                             <td> { imagingCenter.name } </td>
                                             <td>
                                                 <button onClick={ () => this.editImagingCenter(imagingCenter.imagingCenterId)} className="btn btn-outlie-info btn-sm">Update </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.deleteImagingCenter(imagingCenter.imagingCenterId)} className="btn btn-danger btn-sm">Delete </button>
                                                 <button style={{marginLeft: "10px"}} onClick={ () => this.viewImagingCenter(imagingCenter.imagingCenterId)} className="btn btn-outline-info btn-sm">View </button>
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

export default ListImagingCenterComponent
