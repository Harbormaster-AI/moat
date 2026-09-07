
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexLaboratoryOrderComponent } from './index.component';
import { LaboratoryOrderService } from '../../../services/LaboratoryOrder.service';

describe('IndexLaboratoryOrderComponent', () => {
  let component: IndexLaboratoryOrderComponent;
  let fixture: ComponentFixture<IndexLaboratoryOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexLaboratoryOrderComponent
      ],
      providers: [
        LaboratoryOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexLaboratoryOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});