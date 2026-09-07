import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CreativeAssetService } from '../../../services/CreativeAsset.service';
import { SubBaseComponent } from '../../CreativeAsset/sub.base.component';


@Component({
    selector: 'app-edit-creativeAsset',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCreativeAssetComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CreativeAsset';

    creativeAssetForm: FormGroup;
    creativeAsset: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CreativeAssetService,
        private fb: FormBuilder
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

    
    updateCreativeAsset(name, clickUrl, landingPage, width, height, durationSeconds, Files, Approvals, Variations, LineItems, CreativeType, AdFormat): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCreativeAsset(name, clickUrl, landingPage, width, height, durationSeconds, Files, Approvals, Variations, LineItems, CreativeType, AdFormat, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCreativeAsset']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCreativeAsset(params['id']).subscribe(res => {
                this.creativeAsset = res;
            });
        });
    }
}