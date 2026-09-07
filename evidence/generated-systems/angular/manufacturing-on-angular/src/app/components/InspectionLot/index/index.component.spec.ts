
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInspectionLotComponent } from './index.component';
import { InspectionLotService } from '../../../services/InspectionLot.service';

describe('IndexInspectionLotComponent', () => {
  let component: IndexInspectionLotComponent;
  let fixture: ComponentFixture<IndexInspectionLotComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInspectionLotComponent
      ],
      providers: [
        InspectionLotService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInspectionLotComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});