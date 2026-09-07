import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { CreativeVariationService } from '../../../services/CreativeVariation.service';
import { SubBaseComponent } from '../../CreativeVariation/sub.base.component';


@Component({
    selector: 'app-edit-creativeVariation',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditCreativeVariationComponent extends SubBaseComponent implements OnInit {

    title = 'Edit CreativeVariation';

    creativeVariationForm: FormGroup;
    creativeVariation: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: CreativeVariationService,
        private fb: FormBuilder
) {
        super(http);
        this.creativeVariationForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  name: ['', Validators.required],
      language: ['', Validators.required],
      headline: ['', Validators.required],
      bodyText: ['', Validators.required],
      callToAction: ['', Validators.required],
      CreativeAsset: ['', ]
        });
    }

    
    updateCreativeVariation(name, language, headline, bodyText, callToAction, CreativeAsset): void {
        this.route.params.subscribe((params) => {

                        this.service.updateCreativeVariation(name, language, headline, bodyText, callToAction, CreativeAsset, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexCreativeVariation']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getCreativeVariation(params['id']).subscribe(res => {
                this.creativeVariation = res;
            });
        });
    }
}