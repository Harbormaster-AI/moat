import React, { Component } from 'react'
import Component_Service from '../services/Component_Service';

class UpdateComponent_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                partNumber: '',
                name: '',
                componentCategory: '',
                serializationMethod: ''
        }
        this.updateComponent_ = this.updateComponent_.bind(this);

        this.changepartNumberHandler = this.changepartNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeComponentCategoryHandler = this.changeComponentCategoryHandler.bind(this);
        this.changeSerializationMethodHandler = this.changeSerializationMethodHandler.bind(this);
    }

    componentDidMount(){
        Component_Service.getComponent_ById(this.state.id).then( (res) =>{
            let component_ = res.data;
            this.setState({
                partNumber: component_.partNumber,
                name: component_.name,
                componentCategory: component_.componentCategory,
                serializationMethod: component_.serializationMethod
            });
        });
    }

    updateComponent_ = (e) => {
        e.preventDefault();
        let component_ = {
            component_Id: this.state.id,
            partNumber: this.state.partNumber,
            name: this.state.name,
            componentCategory: this.state.componentCategory,
            serializationMethod: this.state.serializationMethod
        };
        console.log('component_ => ' + JSON.stringify(component_));
        console.log('id => ' + JSON.stringify(this.state.id));
        Component_Service.updateComponent_(component_).then( res => {
            this.props.history.push('/component_s');
        });
    }

    changepartNumberHandler= (event) => {
        this.setState({partNumber: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeComponentCategoryHandler= (event) => {
        this.setState({componentCategory: event.target.value});
    }
    changeSerializationMethodHandler= (event) => {
        this.setState({serializationMethod: event.target.value});
    }

    cancel(){
        this.props.history.push('/component_s');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update Component_</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> partNumber: </label>
                                                <input placeholder="partNumber" name="partNumber" className="form-control" value={this.state.partNumber} onChange={this.changepartNumberHandler}/>

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> ComponentCategory: </label>
                                                <select value={this.state.componentCategory} onChange={this.changeComponentCategoryHandler}>
                      <option name="ComponentCategory" className="form-control" >
                          Structure
                      </option>
                      <option name="ComponentCategory" className="form-control" >
                          System
                      </option>
                      <option name="ComponentCategory" className="form-control" >
                          Avionics
                      </option>
                      <option name="ComponentCategory" className="form-control" >
                          Interior
                      </option>
                      <option name="ComponentCategory" className="form-control" >
                          LandingGear
                      </option>
                      <option name="ComponentCategory" className="form-control" >
                          Powerplant
                      </option>
                      <option name="ComponentCategory" className="form-control" >
                          Consumable
                      </option>
                    </select>

                                            <label> SerializationMethod: </label>
                                                <select value={this.state.serializationMethod} onChange={this.changeSerializationMethodHandler}>
                      <option name="SerializationMethod" className="form-control" >
                          Serialized
                      </option>
                      <option name="SerializationMethod" className="form-control" >
                          LotTracked
                      </option>
                      <option name="SerializationMethod" className="form-control" >
                          None
                      </option>
                    </select>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateComponent_}>Save</button>
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

export default UpdateComponent_Component
