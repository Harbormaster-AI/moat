import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CreativeAssetService } from '../../../services/CreativeAsset.service';
import { CreativeAsset } from '../../../models/CreativeAsset';
import { SubBaseComponent } from '../../CreativeAsset/sub.base.component';

@Component({
    selector: 'app-create-creativeAsset',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCreativeAssetComponent extends SubBaseComponent implements OnInit {

    title = 'Add CreativeAsset';

    creativeAssetForm: FormGroup;
    creativeAsset: CreativeAsset;

    constructor( http: HttpClient,
        private creativeAssetService: CreativeAssetService,
        private fb: FormBuilder,
        private router: Router
) {
        super(http);
        this.creativeAssetForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      clickUrl: ['', Validators.required],
      landingPage: ['', Validators.required],
      width: ['', Validators.required],
      height: ['', Validators.required],
      durationSeconds: ['', Validators.required],
      Files: ['', ],
      Approvals: ['', ],
      Variations: ['', ],
      LineItems: ['', ],
      CreativeType: ['', ],
      AdFormat: ['', ]
        });
    }

    
    addCreativeAsset(name, clickUrl, landingPage, width, height, durationSeconds, Files, Approvals, Variations, LineItems, CreativeType, AdFormat): void {
        this.creativeAssetService
        .addCreativeAsset(name, clickUrl, landingPage, width, height, durationSeconds, Files, Approvals, Variations, LineItems, CreativeType, AdFormat)
            .subscribe(() => {
                this.router.navigate(['/indexCreativeAsset']);
            });
    }

    ngOnInit(): void {
    }
}