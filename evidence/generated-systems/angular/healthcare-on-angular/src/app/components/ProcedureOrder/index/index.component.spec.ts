
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexProcedureOrderComponent } from './index.component';
import { ProcedureOrderService } from '../../../services/ProcedureOrder.service';

describe('IndexProcedureOrderComponent', () => {
  let component: IndexProcedureOrderComponent;
  let fixture: ComponentFixture<IndexProcedureOrderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexProcedureOrderComponent
      ],
      providers: [
        ProcedureOrderService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexProcedureOrderComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});