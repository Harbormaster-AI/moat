import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { AudienceSegmentService } from '../../../services/AudienceSegment.service';
import { AudienceSegment } from '../../../models/AudienceSegment';
import { SubBaseComponent } from '../../AudienceSegment/sub.base.component';

@Component({
    selector: 'app-create-audienceSegment',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateAudienceSegmentComponent extends SubBaseComponent implements OnInit {

    title = 'Add AudienceSegment';

    audienceSegmentForm: FormGroup;
    audienceSegment: AudienceSegment;

    constructor( http: HttpClient,
        private audienceSegmentService: AudienceSegmentService,
        private fb: FormBuilder,
        private router: Router
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

    
    addAudienceSegment(name, estimatedReach, description, Provider, Campaigns, ProviderType): void {
        this.audienceSegmentService
        .addAudienceSegment(name, estimatedReach, description, Provider, Campaigns, ProviderType)
            .subscribe(() => {
                this.router.navigate(['/indexAudienceSegment']);
            });
    }

    ngOnInit(): void {
    }
}