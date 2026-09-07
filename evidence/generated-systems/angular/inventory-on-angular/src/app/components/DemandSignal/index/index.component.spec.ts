
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexDemandSignalComponent } from './index.component';
import { DemandSignalService } from '../../../services/DemandSignal.service';

describe('IndexDemandSignalComponent', () => {
  let component: IndexDemandSignalComponent;
  let fixture: ComponentFixture<IndexDemandSignalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexDemandSignalComponent
      ],
      providers: [
        DemandSignalService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexDemandSignalComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});