import React, { Component } from 'react'
import ImagingOrderService from '../services/ImagingOrderService';

class UpdateImagingOrderComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                bodySite: '',
                contrast: '',
                modality: ''
        }
        this.updateImagingOrder = this.updateImagingOrder.bind(this);

        this.changebodySiteHandler = this.changebodySiteHandler.bind(this);
        this.changecontrastHandler = this.changecontrastHandler.bind(this);
        this.changeModalityHandler = this.changeModalityHandler.bind(this);
    }

    componentDidMount(){
        ImagingOrderService.getImagingOrderById(this.state.id).then( (res) =>{
            let imagingOrder = res.data;
            this.setState({
                bodySite: imagingOrder.bodySite,
                contrast: imagingOrder.contrast,
                modality: imagingOrder.modality
            });
        });
    }

    updateImagingOrder = (e) => {
        e.preventDefault();
        let imagingOrder = {
            imagingOrderId: this.state.id,
            bodySite: this.state.bodySite,
            contrast: this.state.contrast,
            modality: this.state.modality
        };
        console.log('imagingOrder => ' + JSON.stringify(imagingOrder));
        console.log('id => ' + JSON.stringify(this.state.id));
        ImagingOrderService.updateImagingOrder(imagingOrder).then( res => {
            this.props.history.push('/imagingOrders');
        });
    }

    changebodySiteHandler= (event) => {
        this.setState({bodySite: event.target.value});
    }
    changecontrastHandler= (event) => {
        this.setState({contrast: event.target.value});
    }
    changeModalityHandler= (event) => {
        this.setState({modality: event.target.value});
    }

    cancel(){
        this.props.history.push('/imagingOrders');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update ImagingOrder</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> bodySite: </label>
                                                <input placeholder="bodySite" name="bodySite" className="form-control" value={this.state.bodySite} onChange={this.changebodySiteHandler}/>

                                            <label> contrast: </label>
                                                <input type="checkbox" placeholder="contrast" name="contrast" className="form-control" value={this.state.contrast} onChange={this.changecontrastHandler}/>


                                            <label> Modality: </label>
                                                <select value={this.state.modality} onChange={this.changeModalityHandler}>
                      <option name="Modality" className="form-control" >
                          XRay
                      </option>
                      <option name="Modality" className="form-control" >
                          CT
                      </option>
                      <option name="Modality" className="form-control" >
                          MRI
                      </option>
                      <option name="Modality" className="form-control" >
                          Ultrasound
                      </option>
                      <option name="Modality" className="form-control" >
                          PET
                      </option>
                      <option name="Modality" className="form-control" >
                          Mammography
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateImagingOrder}>Save</button>
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

export default UpdateImagingOrderComponent
