import React, { Component } from 'react'
import Component_Service from '../services/Component_Service';

class CreateComponent_Component extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                partNumber: '',
                name: '',
                componentCategory: '',
                serializationMethod: ''
        }
        this.changepartNumberHandler = this.changepartNumberHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeComponentCategoryHandler = this.changeComponentCategoryHandler.bind(this);
        this.changeSerializationMethodHandler = this.changeSerializationMethodHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
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
    }
    saveOrUpdateComponent_ = (e) => {
        e.preventDefault();
        let component_ = {
                component_Id: this.state.id,
                partNumber: this.state.partNumber,
                name: this.state.name,
                componentCategory: this.state.componentCategory,
                serializationMethod: this.state.serializationMethod
            };
        console.log('component_ => ' + JSON.stringify(component_));

        // step 5
        if(this.state.id === '_add'){
            component_.component_Id=''
            Component_Service.createComponent_(component_).then(res =>{
                this.props.history.push('/component_s');
            });
        }else{
            Component_Service.updateComponent_(component_).then( res => {
                this.props.history.push('/component_s');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add Component_</h3>
        }else{
            return <h3 className="text-center">Update Component_</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                {
                                    this.getTitle()
                                }
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> partNumber:&emsp; </label>
                                                <input placeholder="partNumber" name="partNumber" className="form-control" value={this.state.partNumber} onChange={this.changepartNumberHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> ComponentCategory:&emsp; </label>
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

                                            <label> SerializationMethod:&emsp; </label>
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

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateComponent_}>Save</button>
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

export default CreateComponent_Component
