import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormGroup, FormBuilder, Validators } from '@angular/forms';

import { BOMItemService } from '../../../services/BOMItem.service';
import { SubBaseComponent } from '../../BOMItem/sub.base.component';


@Component({
    selector: 'app-edit-bOMItem',
    standalone: false,
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.css']
})
export class EditBOMItemComponent extends SubBaseComponent implements OnInit {

    title = 'Edit BOMItem';

    bOMItemForm: FormGroup;
    bOMItem: any;

    constructor( http: HttpClient,
        private route: ActivatedRoute,
        private router: Router,
        private service: BOMItemService,
        private fb: FormBuilder
) {
        super(http);
        this.bOMItemForm = this.createForm();
    }

    createForm(): FormGroup {
        return this.fb.group({
                  lineNumber: ['', Validators.required],
      quantity: ['', Validators.required],
      scrapPercent: ['', Validators.required],
      Bom: ['', ],
      Component: ['', ]
        });
    }

    
    updateBOMItem(lineNumber, quantity, scrapPercent, Bom, Component): void {
        this.route.params.subscribe((params) => {

                        this.service.updateBOMItem(lineNumber, quantity, scrapPercent, Bom, Component, params['id'])
                            .subscribe(() => {
                    this.router.navigate(['/indexBOMItem']);
                });
        });
    }

    ngOnInit(): void {
        this.route.params.subscribe((params) => {
            this.service.getBOMItem(params['id']).subscribe(res => {
                this.bOMItem = res;
            });
        });
    }
}