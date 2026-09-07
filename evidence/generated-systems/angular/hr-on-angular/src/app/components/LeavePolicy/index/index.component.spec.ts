
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexLeavePolicyComponent } from './index.component';
import { LeavePolicyService } from '../../../services/LeavePolicy.service';

describe('IndexLeavePolicyComponent', () => {
  let component: IndexLeavePolicyComponent;
  let fixture: ComponentFixture<IndexLeavePolicyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexLeavePolicyComponent
      ],
      providers: [
        LeavePolicyService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexLeavePolicyComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});