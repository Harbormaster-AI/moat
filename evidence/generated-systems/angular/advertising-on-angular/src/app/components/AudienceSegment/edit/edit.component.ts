import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AudienceSegmentService } from '../../../services/AudienceSegment.service';
import { SubBaseComponent } from '../../AudienceSegment/sub.base.component';


@Component({
    selector: 'app-edit-audienceSegment',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditAudienceSegmentComponent extends SubBaseComponent implements OnInit {

    title = 'Edit AudienceSegment';

    audienceSegmentForm: FormGroup;
    audienceSegment: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: AudienceSegmentService,
        private fb: FormBuilder
) {
        super(http);
        this.audienceSegmentForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      estimatedReach: ['', Validators.required],
      description: ['', Validators.required],
      Provider: ['', ],
      Campaigns: ['', ],
      ProviderType: ['', ]
        });
    }

    
    updateAudienceSegment(name, estimatedReach, description, Provider, Campaigns, ProviderType): void {
        this.route.params.subscribe((params) => {

                        this.service.updateAudienceSegment(name, estimatedReach, description, Provider, Campaigns, ProviderType, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexAudienceSegment']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getAudienceSegment(params['id']).subscribe(res => {
                this.audienceSegment = res;
            });
        });
    }
}