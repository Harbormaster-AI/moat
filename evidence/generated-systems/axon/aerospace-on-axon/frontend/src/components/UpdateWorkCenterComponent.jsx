import React, { Component } from 'react'
import WorkCenterService from '../services/WorkCenterService';

class UpdateWorkCenterComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                capability: ''
        }
        this.updateWorkCenter = this.updateWorkCenter.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changecapabilityHandler = this.changecapabilityHandler.bind(this);
    }

    componentDidMount(){
        WorkCenterService.getWorkCenterById(this.state.id).then( (res) =>{
            let workCenter = res.data;
            this.setState({
                name: workCenter.name,
                capability: workCenter.capability
            });
        });
    }

    updateWorkCenter = (e) => {
        e.preventDefault();
        let workCenter = {
            workCenterId: this.state.id,
            name: this.state.name,
            capability: this.state.capability
        };
        console.log('workCenter => ' + JSON.stringify(workCenter));
        console.log('id => ' + JSON.stringify(this.state.id));
        WorkCenterService.updateWorkCenter(workCenter).then( res => {
            this.props.history.push('/workCenters');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changecapabilityHandler= (event) => {
        this.setState({capability: event.target.value});
    }

    cancel(){
        this.props.history.push('/workCenters');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update WorkCenter</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> capability: </label>
                                                <input placeholder="capability" name="capability" className="form-control" value={this.state.capability} onChange={this.changecapabilityHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateWorkCenter}>Save</button>
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

export default UpdateWorkCenterComponent
