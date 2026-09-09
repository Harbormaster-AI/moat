import React, { Component } from 'react'
import DSPService from '../services/DSPService';

class UpdateDSPComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                website: '',
                region: ''
        }
        this.updateDSP = this.updateDSP.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changewebsiteHandler = this.changewebsiteHandler.bind(this);
        this.changeregionHandler = this.changeregionHandler.bind(this);
    }

    componentDidMount(){
        DSPService.getDSPById(this.state.id).then( (res) =>{
            let dSP = res.data;
            this.setState({
                name: dSP.name,
                website: dSP.website,
                region: dSP.region
            });
        });
    }

    updateDSP = (e) => {
        e.preventDefault();
        let dSP = {
            dSPId: this.state.id,
            name: this.state.name,
            website: this.state.website,
            region: this.state.region
        };
        console.log('dSP => ' + JSON.stringify(dSP));
        console.log('id => ' + JSON.stringify(this.state.id));
        DSPService.updateDSP(dSP).then( res => {
            this.props.history.push('/dSPs');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changewebsiteHandler= (event) => {
        this.setState({website: event.target.value});
    }
    changeregionHandler= (event) => {
        this.setState({region: event.target.value});
    }

    cancel(){
        this.props.history.push('/dSPs');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update DSP</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> website: </label>
                                                <input placeholder="website" name="website" className="form-control" value={this.state.website} onChange={this.changewebsiteHandler}/>

                                            <label> region: </label>
                                                <input placeholder="region" name="region" className="form-control" value={this.state.region} onChange={this.changeregionHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateDSP}>Save</button>
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

export default UpdateDSPComponent
