
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDeviceCriterionComponent } from './index.component';
import { DeviceCriterionService } from '../../../services/DeviceCriterion.service';

describe('IndexDeviceCriterionComponent', () => {
  let component: IndexDeviceCriterionComponent;
  let fixture: ComponentFixture<IndexDeviceCriterionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDeviceCriterionComponent
      ],
      providers: [
        DeviceCriterionService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDeviceCriterionComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});