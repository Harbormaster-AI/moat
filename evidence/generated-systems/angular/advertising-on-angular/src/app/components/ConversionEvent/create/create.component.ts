import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ConversionEventService } from '../../../services/ConversionEvent.service';
import { ConversionEvent } from '../../../models/ConversionEvent';
import { SubBaseComponent } from '../../ConversionEvent/sub.base.component';

@Component({
    selector: 'app-create-conversionEvent',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateConversionEventComponent extends SubBaseComponent implements OnInit {

    title = 'Add ConversionEvent';

    conversionEventForm: FormGroup;
    conversionEvent: ConversionEvent;

    constructor( http: HttpClient,
        private conversionEventService: ConversionEventService,
        private fb: FormBuilder,
        private router: Router
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

    
    addConversionEvent(timestamp, value, Campaign, LineItem, TrackingPixel, EventType, AttributionModel): void {
        this.conversionEventService
        .addConversionEvent(timestamp, value, Campaign, LineItem, TrackingPixel, EventType, AttributionModel)
            .subscribe(() => {
                this.router.navigate(['/indexConversionEvent']);
            });
    }

    ngOnInit(): void {
    }
}