import React, { Component } from 'react'
import AircraftOptionService from '../services/AircraftOptionService';

class CreateAircraftOptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                code: '',
                name: '',
                optionCategory: ''
        }
        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeOptionCategoryHandler = this.changeOptionCategoryHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AircraftOptionService.getAircraftOptionById(this.state.id).then( (res) =>{
                let aircraftOption = res.data;
                this.setState({
                    code: aircraftOption.code,
                    name: aircraftOption.name,
                    optionCategory: aircraftOption.optionCategory
                });
            });
        }        
    }
    saveOrUpdateAircraftOption = (e) => {
        e.preventDefault();
        let aircraftOption = {
                aircraftOptionId: this.state.id,
                code: this.state.code,
                name: this.state.name,
                optionCategory: this.state.optionCategory
            };
        console.log('aircraftOption => ' + JSON.stringify(aircraftOption));

        // step 5
        if(this.state.id === '_add'){
            aircraftOption.aircraftOptionId=''
            AircraftOptionService.createAircraftOption(aircraftOption).then(res =>{
                this.props.history.push('/aircraftOptions');
            });
        }else{
            AircraftOptionService.updateAircraftOption(aircraftOption).then( res => {
                this.props.history.push('/aircraftOptions');
            });
        }
    }
    
    changecodeHandler= (event) => {
        this.setState({code: event.target.value});
    }
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changeOptionCategoryHandler= (event) => {
        this.setState({optionCategory: event.target.value});
    }

    cancel(){
        this.props.history.push('/aircraftOptions');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AircraftOption</h3>
        }else{
            return <h3 className="text-center">Update AircraftOption</h3>
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
                                            <label> code:&emsp; </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> OptionCategory:&emsp; </label>
                                                <select value={this.state.optionCategory} onChange={this.changeOptionCategoryHandler}>
                      <option name="OptionCategory" className="form-control" >
                          Cabin
                      </option>
                      <option name="OptionCategory" className="form-control" >
                          Connectivity
                      </option>
                      <option name="OptionCategory" className="form-control" >
                          Safety
                      </option>
                      <option name="OptionCategory" className="form-control" >
                          Performance
                      </option>
                      <option name="OptionCategory" className="form-control" >
                          Paint
                      </option>
                      <option name="OptionCategory" className="form-control" >
                          FlightDeck
                      </option>
                    </select>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAircraftOption}>Save</button>
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

export default CreateAircraftOptionComponent
