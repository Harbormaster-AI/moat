
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexWorkCenterComponent } from './index.component';
import { WorkCenterService } from '../../../services/WorkCenter.service';

describe('IndexWorkCenterComponent', () => {
  let component: IndexWorkCenterComponent;
  let fixture: ComponentFixture<IndexWorkCenterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexWorkCenterComponent
      ],
      providers: [
        WorkCenterService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexWorkCenterComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});