import React, { Component } from 'react'
import EngineTypeService from '../services/EngineTypeService';

class UpdateEngineTypeComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                engineModelCode: '',
                maxThrustKn: '',
                category: ''
        }
        this.updateEngineType = this.updateEngineType.bind(this);

        this.changeengineModelCodeHandler = this.changeengineModelCodeHandler.bind(this);
        this.changemaxThrustKnHandler = this.changemaxThrustKnHandler.bind(this);
        this.changeCategoryHandler = this.changeCategoryHandler.bind(this);
    }

    componentDidMount(){
        EngineTypeService.getEngineTypeById(this.state.id).then( (res) =>{
            let engineType = res.data;
            this.setState({
                engineModelCode: engineType.engineModelCode,
                maxThrustKn: engineType.maxThrustKn,
                category: engineType.category
            });
        });
    }

    updateEngineType = (e) => {
        e.preventDefault();
        let engineType = {
            engineTypeId: this.state.id,
            engineModelCode: this.state.engineModelCode,
            maxThrustKn: this.state.maxThrustKn,
            category: this.state.category
        };
        console.log('engineType => ' + JSON.stringify(engineType));
        console.log('id => ' + JSON.stringify(this.state.id));
        EngineTypeService.updateEngineType(engineType).then( res => {
            this.props.history.push('/engineTypes');
        });
    }

    changeengineModelCodeHandler= (event) => {
        this.setState({engineModelCode: event.target.value});
    }
    changemaxThrustKnHandler= (event) => {
        this.setState({maxThrustKn: event.target.value});
    }
    changeCategoryHandler= (event) => {
        this.setState({category: event.target.value});
    }

    cancel(){
        this.props.history.push('/engineTypes');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update EngineType</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> engineModelCode: </label>
                                                <input placeholder="engineModelCode" name="engineModelCode" className="form-control" value={this.state.engineModelCode} onChange={this.changeengineModelCodeHandler}/>

                                            <label> maxThrustKn: </label>
                                                <input placeholder="maxThrustKn" name="maxThrustKn" className="form-control" value={this.state.maxThrustKn} onChange={this.changemaxThrustKnHandler}/>

                                            <label> Category: </label>
                                                <select value={this.state.category} onChange={this.changeCategoryHandler}>
                      <option name="Category" className="form-control" >
                          Turbofan
                      </option>
                      <option name="Category" className="form-control" >
                          Turboprop
                      </option>
                      <option name="Category" className="form-control" >
                          Turbojet
                      </option>
                      <option name="Category" className="form-control" >
                          Piston
                      </option>
                      <option name="Category" className="form-control" >
                          Electric
                      </option>
                      <option name="Category" className="form-control" >
                          Rocket
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateEngineType}>Save</button>
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

export default UpdateEngineTypeComponent
