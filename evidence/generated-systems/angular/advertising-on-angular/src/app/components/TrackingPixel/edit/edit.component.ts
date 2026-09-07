import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { TrackingPixelService } from '../../../services/TrackingPixel.service';
import { SubBaseComponent } from '../../TrackingPixel/sub.base.component';


@Component({
    selector: 'app-edit-trackingPixel',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditTrackingPixelComponent extends SubBaseComponent implements OnInit {

    title = 'Edit TrackingPixel';

    trackingPixelForm: FormGroup;
    trackingPixel: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: TrackingPixelService,
        private fb: FormBuilder
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

    
    updateTrackingPixel(name, url, Campaign, Advertiser, ConversionEvents, EventType, PixelType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateTrackingPixel(name, url, Campaign, Advertiser, ConversionEvents, EventType, PixelType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexTrackingPixel']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getTrackingPixel(params['id']).subscribe(res => {
                this.trackingPixel = res;
            });
        });
    }
}