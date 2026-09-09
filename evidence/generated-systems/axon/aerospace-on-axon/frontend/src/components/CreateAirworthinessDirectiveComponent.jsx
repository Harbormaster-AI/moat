import React, { Component } from 'react'
import AirworthinessDirectiveService from '../services/AirworthinessDirectiveService';

class CreateAirworthinessDirectiveComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                directiveNumber: '',
                title: ''
        }
        this.changedirectiveNumberHandler = this.changedirectiveNumberHandler.bind(this);
        this.changetitleHandler = this.changetitleHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            AirworthinessDirectiveService.getAirworthinessDirectiveById(this.state.id).then( (res) =>{
                let airworthinessDirective = res.data;
                this.setState({
                    directiveNumber: airworthinessDirective.directiveNumber,
                    title: airworthinessDirective.title
                });
            });
        }        
    }
    saveOrUpdateAirworthinessDirective = (e) => {
        e.preventDefault();
        let airworthinessDirective = {
                airworthinessDirectiveId: this.state.id,
                directiveNumber: this.state.directiveNumber,
                title: this.state.title
            };
        console.log('airworthinessDirective => ' + JSON.stringify(airworthinessDirective));

        // step 5
        if(this.state.id === '_add'){
            airworthinessDirective.airworthinessDirectiveId=''
            AirworthinessDirectiveService.createAirworthinessDirective(airworthinessDirective).then(res =>{
                this.props.history.push('/airworthinessDirectives');
            });
        }else{
            AirworthinessDirectiveService.updateAirworthinessDirective(airworthinessDirective).then( res => {
                this.props.history.push('/airworthinessDirectives');
            });
        }
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

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add AirworthinessDirective</h3>
        }else{
            return <h3 className="text-center">Update AirworthinessDirective</h3>
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
                                            <label> directiveNumber:&emsp; </label>
                                                <input placeholder="directiveNumber" name="directiveNumber" className="form-control" value={this.state.directiveNumber} onChange={this.changedirectiveNumberHandler}/>

                                            <label> title:&emsp; </label>
                                                <input placeholder="title" name="title" className="form-control" value={this.state.title} onChange={this.changetitleHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateAirworthinessDirective}>Save</button>
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

export default CreateAirworthinessDirectiveComponent
