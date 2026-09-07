
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexTaxWithholdingComponent } from './index.component';
import { TaxWithholdingService } from '../../../services/TaxWithholding.service';

describe('IndexTaxWithholdingComponent', () => {
  let component: IndexTaxWithholdingComponent;
  let fixture: ComponentFixture<IndexTaxWithholdingComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexTaxWithholdingComponent
      ],
      providers: [
        TaxWithholdingService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexTaxWithholdingComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});