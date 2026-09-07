
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexWorkShiftComponent } from './index.component';
import { WorkShiftService } from '../../../services/WorkShift.service';

describe('IndexWorkShiftComponent', () => {
  let component: IndexWorkShiftComponent;
  let fixture: ComponentFixture<IndexWorkShiftComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexWorkShiftComponent
      ],
      providers: [
        WorkShiftService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexWorkShiftComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});