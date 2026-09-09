import React, { Component } from 'react'
import FeatureSetService from '../services/FeatureSetService';

class UpdateFeatureSetComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                refreshSchedule: '',
                storeType: ''
        }
        this.updateFeatureSet = this.updateFeatureSet.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changerefreshScheduleHandler = this.changerefreshScheduleHandler.bind(this);
        this.changeStoreTypeHandler = this.changeStoreTypeHandler.bind(this);
    }

    componentDidMount(){
        FeatureSetService.getFeatureSetById(this.state.id).then( (res) =>{
            let featureSet = res.data;
            this.setState({
                name: featureSet.name,
                refreshSchedule: featureSet.refreshSchedule,
                storeType: featureSet.storeType
            });
        });
    }

    updateFeatureSet = (e) => {
        e.preventDefault();
        let featureSet = {
            featureSetId: this.state.id,
            name: this.state.name,
            refreshSchedule: this.state.refreshSchedule,
            storeType: this.state.storeType
        };
        console.log('featureSet => ' + JSON.stringify(featureSet));
        console.log('id => ' + JSON.stringify(this.state.id));
        FeatureSetService.updateFeatureSet(featureSet).then( res => {
            this.props.history.push('/featureSets');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changerefreshScheduleHandler= (event) => {
        this.setState({refreshSchedule: event.target.value});
    }
    changeStoreTypeHandler= (event) => {
        this.setState({storeType: event.target.value});
    }

    cancel(){
        this.props.history.push('/featureSets');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update FeatureSet</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> refreshSchedule: </label>
                                                <input placeholder="refreshSchedule" name="refreshSchedule" className="form-control" value={this.state.refreshSchedule} onChange={this.changerefreshScheduleHandler}/>

                                            <label> StoreType: </label>
                                                <select value={this.state.storeType} onChange={this.changeStoreTypeHandler}>
                      <option name="StoreType" className="form-control" >
                          Online
                      </option>
                      <option name="StoreType" className="form-control" >
                          Offline
                      </option>
                      <option name="StoreType" className="form-control" >
                          Hybrid
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateFeatureSet}>Save</button>
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

export default UpdateFeatureSetComponent
