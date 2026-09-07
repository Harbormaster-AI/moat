import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { ConversionEventService } from '../../../services/ConversionEvent.service';
import { SubBaseComponent } from '../../ConversionEvent/sub.base.component';


@Component({
    selector: 'app-edit-conversionEvent',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditConversionEventComponent extends SubBaseComponent implements OnInit {

    title = 'Edit ConversionEvent';

    conversionEventForm: FormGroup;
    conversionEvent: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: ConversionEventService,
        private fb: FormBuilder
) {
        super(http);
        this.conversionEventForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  timestamp: ['', Validators.required],
      value: ['', Validators.required],
      Campaign: ['', ],
      LineItem: ['', ],
      TrackingPixel: ['', ],
      EventType: ['', ],
      AttributionModel: ['', ]
        });
    }

    
    updateConversionEvent(timestamp, value, Campaign, LineItem, TrackingPixel, EventType, AttributionModel): void {
        this.route.params.subscribe((params) => {

                        this.service.updateConversionEvent(timestamp, value, Campaign, LineItem, TrackingPixel, EventType, AttributionModel, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexConversionEvent']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getConversionEvent(params['id']).subscribe(res => {
                this.conversionEvent = res;
            });
        });
    }
}