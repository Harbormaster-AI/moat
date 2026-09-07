
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { IndexStorageLocationComponent } from './index.component';
import { StorageLocationService } from '../../../services/StorageLocation.service';

describe('IndexStorageLocationComponent', () => {
  let component: IndexStorageLocationComponent;
  let fixture: ComponentFixture<IndexStorageLocationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        IndexStorageLocationComponent
      ],
      providers: [
        StorageLocationService,
        {
          provide: Router,
          useValue: {
            navigate: jasmine.createSpy('navigate'),
            navigateByUrl: jasmine.createSpy('navigateByUrl')
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(IndexStorageLocationComponent);
    component = fixture.componentInstance;

    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});