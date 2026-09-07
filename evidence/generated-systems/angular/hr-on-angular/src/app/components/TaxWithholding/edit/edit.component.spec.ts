
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { EditTaxWithholdingComponent } from './edit.component';
import { TaxWithholdingService } from '../../../services/TaxWithholding.service';

describe('EditTaxWithholdingComponent', () => {
  let component: EditTaxWithholdingComponent;
  let fixture: ComponentFixture<EditTaxWithholdingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        ReactiveFormsModule
      ],
      declarations: [
        EditTaxWithholdingComponent
      ],
      providers: [
        TaxWithholdingService,
        {
          provide: ActivatedRoute,
          useValue: {
            params: of({ id: '1' })
          }
        },
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(EditTaxWithholdingComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});