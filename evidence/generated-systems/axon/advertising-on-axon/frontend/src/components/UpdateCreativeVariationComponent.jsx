import React, { Component } from 'react'
import CreativeVariationService from '../services/CreativeVariationService';

class UpdateCreativeVariationComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            id: this.props.match.params.id,
                name: '',
                language: '',
                headline: '',
                bodyText: '',
                callToAction: ''
        }
        this.updateCreativeVariation = this.updateCreativeVariation.bind(this);

        this.changenameHandler = this.changenameHandler.bind(this);
        this.changelanguageHandler = this.changelanguageHandler.bind(this);
        this.changeheadlineHandler = this.changeheadlineHandler.bind(this);
        this.changebodyTextHandler = this.changebodyTextHandler.bind(this);
        this.changecallToActionHandler = this.changecallToActionHandler.bind(this);
    }

    componentDidMount(){
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

    updateCreativeVariation = (e) => {
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
        console.log('id => ' + JSON.stringify(this.state.id));
        CreativeVariationService.updateCreativeVariation(creativeVariation).then( res => {
            this.props.history.push('/creativeVariations');
        });
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

    render() {
        return (
            <div>
                <br></br>
                   <div className = "container">
                        <div className = "row">
                            <div className = "card col-md-6 offset-md-3 offset-md-3">
                                <h3 className="text-center">Update CreativeVariation</h3>
                                <div className = "card-body">
                                    <form>
                                        <div className = "form-group">
                                            <label> name: </label>
                                                <input placeholder="name" name="name" className="form-control" value={this.state.name} onChange={this.changenameHandler}/>

                                            <label> language: </label>
                                                <input placeholder="language" name="language" className="form-control" value={this.state.language} onChange={this.changelanguageHandler}/>

                                            <label> headline: </label>
                                                <input placeholder="headline" name="headline" className="form-control" value={this.state.headline} onChange={this.changeheadlineHandler}/>

                                            <label> bodyText: </label>
                                                <input placeholder="bodyText" name="bodyText" className="form-control" value={this.state.bodyText} onChange={this.changebodyTextHandler}/>

                                            <label> callToAction: </label>
                                                <input placeholder="callToAction" name="callToAction" className="form-control" value={this.state.callToAction} onChange={this.changecallToActionHandler}/>

                                        </div>
                                        <button className="btn btn-success" onClick={this.updateCreativeVariation}>Save</button>
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

export default UpdateCreativeVariationComponent
