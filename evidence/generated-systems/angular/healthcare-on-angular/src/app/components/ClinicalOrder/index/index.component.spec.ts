
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexClinicalOrderComponent } from './index.component';
import { ClinicalOrderService } from '../../../services/ClinicalOrder.service';

describe('IndexClinicalOrderComponent', () => {
  let component: IndexClinicalOrderComponent;
  let fixture: ComponentFixture<IndexClinicalOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexClinicalOrderComponent
      ],
      providers: [
        ClinicalOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexClinicalOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});