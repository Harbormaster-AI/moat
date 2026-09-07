
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexSerialNumberComponent } from './index.component';
import { SerialNumberService } from '../../../services/SerialNumber.service';

describe('IndexSerialNumberComponent', () => {
  let component: IndexSerialNumberComponent;
  let fixture: ComponentFixture<IndexSerialNumberComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexSerialNumberComponent
      ],
      providers: [
        SerialNumberService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexSerialNumberComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});