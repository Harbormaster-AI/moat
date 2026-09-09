import React, { Component } from 'react'
import ImagingCenterService from '../services/ImagingCenterService';

class UpdateImagingCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: ''
        }
        this.updateImagingCenter = this.updateImagingCenter.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
    }

    componentDidMount(){
        ImagingCenterService.getImagingCenterById(this.state.id).then( (res) =>{
            let imagingCenter = res.data;
            this.setState({
                name: imagingCenter.name
            });
        });
    }

    updateImagingCenter = (e) => {
        e.preventDefault();
        let imagingCenter = {
            imagingCenterId: this.state.id,
            name: this.state.name
        };
        console.log('imagingCenter => ' + JSON.stringify(imagingCenter));
        console.log('id => ' + JSON.stringify(this.state.id));
        ImagingCenterService.updateImagingCenter(imagingCenter).then( res => {
            this.props.history.push('/imagingCenters');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }

    cancel(){
        this.props.history.push('/imagingCenters');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ImagingCenter</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateImagingCenter}>Save</button>
                                        <button className="btn btn-danger" onClick={this.cancel.bind(this)} style={{marginLeft: "10px"}}>Cancel</button>
                                    </form>
                                </div>
                            </div>
                        </div>

                   </div>
            </div>
        )
    }
}

export default UpdateImagingCenterComponent
