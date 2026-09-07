import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { TrackingPixelService } from '../../../services/TrackingPixel.service';
import { TrackingPixel } from '../../../models/TrackingPixel';
import { SubBaseComponent } from '../../TrackingPixel/sub.base.component';

@Component({
    selector: 'app-create-trackingPixel',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateTrackingPixelComponent extends SubBaseComponent implements OnInit {

    title = 'Add TrackingPixel';

    trackingPixelForm: FormGroup;
    trackingPixel: TrackingPixel;

    constructor( http: HttpClient,
        private trackingPixelService: TrackingPixelService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.trackingPixelForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      url: ['', Validators.required],
      Campaign: ['', ],
      Advertiser: ['', ],
      ConversionEvents: ['', ],
      EventType: ['', ],
      PixelType: ['', ]
        });
    }

    
    addTrackingPixel(name, url, Campaign, Advertiser, ConversionEvents, EventType, PixelType): void {
        this.trackingPixelService
        .addTrackingPixel(name, url, Campaign, Advertiser, ConversionEvents, EventType, PixelType)
            .subscribe(() => {
                this.router.navigate(['/indexTrackingPixel']);
            });
    }

    ngOnInit(): void {
    }
}