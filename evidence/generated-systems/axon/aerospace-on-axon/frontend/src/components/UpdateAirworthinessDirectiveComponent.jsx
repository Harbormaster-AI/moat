import React, { Component } from 'react'
import AirworthinessDirectiveService from '../services/AirworthinessDirectiveService';

class UpdateAirworthinessDirectiveComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                directiveNumber: '',
                title: ''
        }
        this.updateAirworthinessDirective = this.updateAirworthinessDirective.bind(this);

        this.changedirectiveNumberHandler = this.changedirectiveNumberHandler.bind(this);
        this.changetitleHandler = this.changetitleHandler.bind(this);
    }

    componentDidMount(){
        AirworthinessDirectiveService.getAirworthinessDirectiveById(this.state.id).then( (res) =>{
            let airworthinessDirective = res.data;
            this.setState({
                directiveNumber: airworthinessDirective.directiveNumber,
                title: airworthinessDirective.title
            });
        });
    }

    updateAirworthinessDirective = (e) => {
        e.preventDefault();
        let airworthinessDirective = {
            airworthinessDirectiveId: this.state.id,
            directiveNumber: this.state.directiveNumber,
            title: this.state.title
        };
        console.log('airworthinessDirective => ' + JSON.stringify(airworthinessDirective));
        console.log('id => ' + JSON.stringify(this.state.id));
        AirworthinessDirectiveService.updateAirworthinessDirective(airworthinessDirective).then( res => {
            this.props.history.push('/airworthinessDirectives');
        });
    }

    changedirectiveNumberHandler= (event) => {
        this.setState({directiveNumber: event.target.value});
    }
    changetitleHandler= (event) => {
        this.setState({title: event.target.value});
    }

    cancel(){
        this.props.history.push('/airworthinessDirectives');
    }

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update AirworthinessDirective</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> directiveNumber: </label>
                                                <input placeholder="directiveNumber" name="directiveNumber" className="form-control" value={this.state.directiveNumber} onChange={this.changedirectiveNumberHandler}/>

                                            <label> title: </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateAirworthinessDirective}>Save</button>
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

export default UpdateAirworthinessDirectiveComponent
