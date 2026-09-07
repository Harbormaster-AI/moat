
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexInspectionCharacteristicComponent } from './index.component';
import { InspectionCharacteristicService } from '../../../services/InspectionCharacteristic.service';

describe('IndexInspectionCharacteristicComponent', () => {
  let component: IndexInspectionCharacteristicComponent;
  let fixture: ComponentFixture<IndexInspectionCharacteristicComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexInspectionCharacteristicComponent
      ],
      providers: [
        InspectionCharacteristicService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexInspectionCharacteristicComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});