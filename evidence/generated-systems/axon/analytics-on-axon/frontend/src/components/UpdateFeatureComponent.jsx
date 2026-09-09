import React, { Component } from 'react'
import FeatureService from '../services/FeatureService';

class UpdateFeatureComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                description: '',
                dataType: ''
        }
        this.updateFeature = this.updateFeature.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changedescriptionHandler = this.changedescriptionHandler.bind(this);
        this.changeDataTypeHandler = this.changeDataTypeHandler.bind(this);
    }

    componentDidMount(){
        FeatureService.getFeatureById(this.state.id).then( (res) =>{
            let feature = res.data;
            this.setState({
                name: feature.name,
                description: feature.description,
                dataType: feature.dataType
            });
        });
    }

    updateFeature = (e) => {
        e.preventDefault();
        let feature = {
            featureId: this.state.id,
            name: this.state.name,
            description: this.state.description,
            dataType: this.state.dataType
        };
        console.log('feature => ' + JSON.stringify(feature));
        console.log('id => ' + JSON.stringify(this.state.id));
        FeatureService.updateFeature(feature).then( res => {
            this.props.history.push('/features');
        });
    }

    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changedescriptionHandler= (event) => {
        this.setState({description: event.target.value});
    }
    changeDataTypeHandler= (event) => {
        this.setState({dataType: event.target.value});
    }

    cancel(){
        this.props.history.push('/features');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Feature</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> description: </label>
                                                <input placeholder="description" name="description" className="form-control" value={this.state.description} onChange={this.changedescriptionHandler}/>

                                            <label> DataType: </label>
                                                <select value={this.state.dataType} onChange={this.changeDataTypeHandler}>
                      <option name="DataType" className="form-control" >
                          String
                      </option>
                      <option name="DataType" className="form-control" >
                          Integer
                      </option>
                      <option name="DataType" className="form-control" >
                          Decimal
                      </option>
                      <option name="DataType" className="form-control" >
                          Boolean
                      </option>
                      <option name="DataType" className="form-control" >
                          Date
                      </option>
                      <option name="DataType" className="form-control" >
                          DateTime
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateFeature}>Save</button>
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

export default UpdateFeatureComponent
