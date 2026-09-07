
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexPayrollItemComponent } from './index.component';
import { PayrollItemService } from '../../../services/PayrollItem.service';

describe('IndexPayrollItemComponent', () => {
  let component: IndexPayrollItemComponent;
  let fixture: ComponentFixture<IndexPayrollItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexPayrollItemComponent
      ],
      providers: [
        PayrollItemService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexPayrollItemComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});