import React, { Component } from 'react'
import CreativeVariationService from '../services/CreativeVariationService';

class CreateCreativeVariationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
                name: '',
                language: '',
                headline: '',
                bodyText: '',
                callToAction: ''
        }
        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelanguageHandler = this.changelanguageHandler.bind(this);
        this.changeheadlineHandler = this.changeheadlineHandler.bind(this);
        this.changebodyTextHandler = this.changebodyTextHandler.bind(this);
        this.changecallToActionHandler = this.changecallToActionHandler.bind(this);
    }

    // step 3
    componentDidMount(){

        // step 4
        if(this.state.id === '_add'){
            return
        }else{
            CreativeVariationService.getCreativeVariationById(this.state.id).then( (res) =>{
                let creativeVariation = res.data;
                this.setState({
                    name: creativeVariation.name,
                    language: creativeVariation.language,
                    headline: creativeVariation.headline,
                    bodyText: creativeVariation.bodyText,
                    callToAction: creativeVariation.callToAction
                });
            });
        }        
    }
    saveOrUpdateCreativeVariation = (e) => {
        e.preventDefault();
        let creativeVariation = {
                creativeVariationId: this.state.id,
                name: this.state.name,
                language: this.state.language,
                headline: this.state.headline,
                bodyText: this.state.bodyText,
                callToAction: this.state.callToAction
            };
        console.log('creativeVariation => ' + JSON.stringify(creativeVariation));

        // step 5
        if(this.state.id === '_add'){
            creativeVariation.creativeVariationId=''
            CreativeVariationService.createCreativeVariation(creativeVariation).then(res =>{
                this.props.history.push('/creativeVariations');
            });
        }else{
            CreativeVariationService.updateCreativeVariation(creativeVariation).then( res => {
                this.props.history.push('/creativeVariations');
            });
        }
    }
    
    changenameHandler= (event) => {
        this.setState({name: event.target.value});
    }
    changelanguageHandler= (event) => {
        this.setState({language: event.target.value});
    }
    changeheadlineHandler= (event) => {
        this.setState({headline: event.target.value});
    }
    changebodyTextHandler= (event) => {
        this.setState({bodyText: event.target.value});
    }
    changecallToActionHandler= (event) => {
        this.setState({callToAction: event.target.value});
    }

    cancel(){
        this.props.history.push('/creativeVariations');
    }

    getTitle(){
        if(this.state.id === '_add'){
            return <h3 className="text-center">Add CreativeVariation</h3>
        }else{
            return <h3 className="text-center">Update CreativeVariation</h3>
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
                                            <label> name:&emsp; </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> language:&emsp; </label>
                                                <input placeholder="language" name="language" className="form-control" value={this.state.language} onChange={this.changelanguageHandler}/>

                                            <label> headline:&emsp; </label>
                                                <input placeholder="headline" name="headline" className="form-control" value={this.state.headline} onChange={this.changeheadlineHandler}/>

                                            <label> bodyText:&emsp; </label>
                                                <input placeholder="bodyText" name="bodyText" className="form-control" value={this.state.bodyText} onChange={this.changebodyTextHandler}/>

                                            <label> callToAction:&emsp; </label>
                                                <input placeholder="callToAction" name="callToAction" className="form-control" value={this.state.callToAction} onChange={this.changecallToActionHandler}/>

                                        </div>

                                        <button className="btn btn-outline-success" onClick={this.saveOrUpdateCreativeVariation}>Save</button>
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

export default CreateCreativeVariationComponent
