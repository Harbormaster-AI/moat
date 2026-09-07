import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CreativeVariationService } from '../../../services/CreativeVariation.service';
import { CreativeVariation } from '../../../models/CreativeVariation';
import { SubBaseComponent } from '../../CreativeVariation/sub.base.component';

@Component({
    selector: 'app-create-creativeVariation',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateCreativeVariationComponent extends SubBaseComponent implements OnInit {

    title = 'Add CreativeVariation';

    creativeVariationForm: FormGroup;
    creativeVariation: CreativeVariation;

    constructor( http: HttpClient,
        private creativeVariationService: CreativeVariationService,
        private fb: FormBuilder,
        private router: Router
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

    
    addCreativeVariation(name, language, headline, bodyText, callToAction, CreativeAsset): void {
        this.creativeVariationService
        .addCreativeVariation(name, language, headline, bodyText, callToAction, CreativeAsset)
            .subscribe(() => {
                this.router.navigate(['/indexCreativeVariation']);
            });
    }

    ngOnInit(): void {
    }
}