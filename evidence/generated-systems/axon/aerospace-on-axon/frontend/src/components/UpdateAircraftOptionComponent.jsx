import React, { Component } from 'react'
import AircraftOptionService from '../services/AircraftOptionService';

class UpdateAircraftOptionComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                code: '',
                name: '',
                optionCategory: ''
        }
        this.updateAircraftOption = this.updateAircraftOption.bind(this);

        this.changecodeHandler = this.changecodeHandler.bind(this);
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changeOptionCategoryHandler = this.changeOptionCategoryHandler.bind(this);
    }

    componentDidMount(){
        AircraftOptionService.getAircraftOptionById(this.state.id).then( (res) =>{
            let aircraftOption = res.data;
            this.setState({
                code: aircraftOption.code,
                name: aircraftOption.name,
                optionCategory: aircraftOption.optionCategory
            });
        });
    }

    updateAircraftOption = (e) => {
        e.preventDefault();
        let aircraftOption = {
            aircraftOptionId: this.state.id,
            code: this.state.code,
            name: this.state.name,
            optionCategory: this.state.optionCategory
        };
        console.log('aircraftOption => ' + JSON.stringify(aircraftOption));
        console.log('id => ' + JSON.stringify(this.state.id));
        AircraftOptionService.updateAircraftOption(aircraftOption).then( res => {
            this.props.history.push('/aircraftOptions');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AircraftOption</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> code: </label>
                                                <input placeholder="code" name="code" className="form-control" value={this.state.code} onChange={this.changecodeHandler}/>

                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> OptionCategory: </label>
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
                                        <button className="btn btn-success" onClick={this.updateAircraftOption}>Save</button>
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

export default UpdateAircraftOptionComponent
