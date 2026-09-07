import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { BOMItemService } from '../../../services/BOMItem.service';
import { BOMItem } from '../../../models/BOMItem';
import { SubBaseComponent } from '../../BOMItem/sub.base.component';

@Component({
    selector: 'app-create-bOMItem',
    standalone: false,
    templateUrl: './create.component.html',
    styleUrls: ['./create.component.css']
})
export class CreateBOMItemComponent extends SubBaseComponent implements OnInit {

    title = 'Add BOMItem';

    bOMItemForm: FormGroup;
    bOMItem: BOMItem;

    constructor( http: HttpClient,
        private bOMItemService: BOMItemService,
        private fb: FormBuilder,
        private router: Router
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

    
    addBOMItem(lineNumber, quantity, scrapPercent, Bom, Component): void {
        this.bOMItemService
        .addBOMItem(lineNumber, quantity, scrapPercent, Bom, Component)
            .subscribe(() => {
                this.router.navigate(['/indexBOMItem']);
            });
    }

    ngOnInit(): void {
    }
}