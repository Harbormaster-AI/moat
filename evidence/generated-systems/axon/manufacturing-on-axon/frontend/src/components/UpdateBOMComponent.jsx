import React, { Component } from 'react'
import BOMService from '../services/BOMService';

class UpdateBOMComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                bomNumber: '',
                revision: '',
                effectivityStart: '',
                effectivityEnd: '',
                status: ''
        }
        this.updateBOM = this.updateBOM.bind(this);

        this.changebomNumberHandler = this.changebomNumberHandler.bind(this);
        this.changerevisionHandler = this.changerevisionHandler.bind(this);
        this.changeeffectivityStartHandler = this.changeeffectivityStartHandler.bind(this);
        this.changeeffectivityEndHandler = this.changeeffectivityEndHandler.bind(this);
        this.changeStatusHandler = this.changeStatusHandler.bind(this);
    }

    componentDidMount(){
        BOMService.getBOMById(this.state.id).then( (res) =>{
            let bOM = res.data;
            this.setState({
                bomNumber: bOM.bomNumber,
                revision: bOM.revision,
                effectivityStart: bOM.effectivityStart,
                effectivityEnd: bOM.effectivityEnd,
                status: bOM.status
            });
        });
    }

    updateBOM = (e) => {
        e.preventDefault();
        let bOM = {
            bOMId: this.state.id,
            bomNumber: this.state.bomNumber,
            revision: this.state.revision,
            effectivityStart: this.state.effectivityStart,
            effectivityEnd: this.state.effectivityEnd,
            status: this.state.status
        };
        console.log('bOM => ' + JSON.stringify(bOM));
        console.log('id => ' + JSON.stringify(this.state.id));
        BOMService.updateBOM(bOM).then( res => {
            this.props.history.push('/bOMs');
        });
    }

    changebomNumberHandler= (event) => {
        this.setState({bomNumber: event.target.value});
    }
    changerevisionHandler= (event) => {
        this.setState({revision: event.target.value});
    }
    changeeffectivityStartHandler= (event) => {
        this.setState({effectivityStart: event.target.value});
    }
    changeeffectivityEndHandler= (event) => {
        this.setState({effectivityEnd: event.target.value});
    }
    changeStatusHandler= (event) => {
        this.setState({status: event.target.value});
    }

    cancel(){
        this.props.history.push('/bOMs');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update BOM</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> bomNumber: </label>
                                                <input placeholder="bomNumber" name="bomNumber" className="form-control" value={this.state.bomNumber} onChange={this.changebomNumberHandler}/>

                                            <label> revision: </label>
                                                <input placeholder="revision" name="revision" className="form-control" value={this.state.revision} onChange={this.changerevisionHandler}/>

                                            <label> effectivityStart: </label>
                                                <input type="date" placeholder="effectivityStart" name="effectivityStart" className="form-control" value={this.state.effectivityStart} onChange={this.changeeffectivityStartHandler}/>

                                            <label> effectivityEnd: </label>
                                                <input type="date" placeholder="effectivityEnd" name="effectivityEnd" className="form-control" value={this.state.effectivityEnd} onChange={this.changeeffectivityEndHandler}/>

                                            <label> Status: </label>
                                                <select value={this.state.status} onChange={this.changeStatusHandler}>
                      <option name="Status" className="form-control" >
                          Draft
                      </option>
                      <option name="Status" className="form-control" >
                          Released
                      </option>
                      <option name="Status" className="form-control" >
                          Obsolete
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateBOM}>Save</button>
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

export default UpdateBOMComponent
